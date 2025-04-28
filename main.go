package main

import (
	"context"
	"fmt"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/sqlca/v2"
	"os"
	"os/signal"

	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/logic"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/godotenv"
	"github.com/civet148/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

const (
	Version     = "0.1.1"
	ProgramName = "gocex"
)

var (
	BuildTime = "2025-04-25"
	GitCommit = "<N/A>"
)

const (
	CmdName_Pos   = "pos"   //仓位子命令
	CmdName_Open  = "open"  //建仓
	CmdName_Close = "close" //平仓
	CmdName_List  = "list"  //仓位列表
)

const (
	CmdFlag_Config      = "config"
	CmdFlag_Debug       = "debug"
	CmdFlag_Cex         = "cex"
	CmdFlag_Symbol      = "symbol"
	CmdFlag_Px          = "px"
	CmdFlag_Sz          = "sz"
	CmdFlag_Lever       = "lever"
	CmdFlag_OrderType   = "order-type"
	CmdFlag_PosSideType = "pos-side"
	CmdFlag_TradeMode   = "trade-mode"
	CmdFlag_OrderId     = "order-id"
)

func init() {
	log.SetLevel("debug")
	err := log.Open("gocex.log")
	if err != nil {
		panic(err.Error())
	}
}

func grace() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	go func() {
		for {
			select {
			case s := <-sigChannel:
				{
					if s != nil && s == os.Interrupt {
						fmt.Printf("Ctrl+C signal captured, program exiting...\n")
						close(sigChannel)
						os.Exit(0)
					}
				}
			}
		}
	}()
}

func loadConfig(ctx *cli.Context) (*config.Config, error) {
	var c config.Config

	//设置配置文件
	viper.SetConfigFile(ctx.String(CmdFlag_Config))

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		return nil, log.Errorf(err)
	}

	// 反序列化到结构体
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, log.Errorf(err)
	}
	if c.ApiKey == "" {
		err = godotenv.Load(&c)
		if err != nil {
			log.Errorf("load .env error %s", err)
			return nil, err
		}
	}

	log.Json(&c)
	return &c, nil
}

func main() {

	grace()

	app := &cli.App{
		Name:    ProgramName,
		Usage:   "",
		Version: fmt.Sprintf("v%s %s commit %s", Version, BuildTime, GitCommit),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    CmdFlag_Config,
				Usage:   "config file",
				Aliases: []string{"c"},
				Value:   "config.yaml",
			},
			&cli.BoolFlag{
				Name:    CmdFlag_Debug,
				Usage:   "debug",
				Aliases: []string{"d"},
			},
		},
		Commands: []*cli.Command{
			cmdPos,
		},
		Action: func(ctx *cli.Context) error {

			if ctx.Bool(CmdFlag_Debug) {
				log.SetLevel("debug")
			}
			var err error
			var c *config.Config
			c, err = loadConfig(ctx)
			if err != nil {
				log.Panic(err.Error())
			}
			if c.Debug {
				log.SetLevel("debug")
			}

			cex := logic.NewCexLogic(c)
			return cex.Run()
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

var cmdPos = &cli.Command{
	Name:    CmdName_Pos,
	Aliases: []string{"p"},
	Usage:   "position command",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
	},
	Subcommands: []*cli.Command{
		cmdPosOpen,
		cmdPosClose,
		cmdPosList,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

var cmdPosOpen = &cli.Command{
	Name:    CmdName_Open,
	Aliases: []string{"o"},
	Usage:   "create position",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
		&cli.StringFlag{
			Name:    CmdFlag_Symbol,
			Aliases: []string{"s"},
			Value:   types.PEPEUSDT,
		},
		&cli.StringFlag{
			Name:    CmdFlag_Px,
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:     CmdFlag_Sz,
			Aliases:  []string{"z"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    CmdFlag_Lever,
			Aliases: []string{"l"},
			Value:   "1",
		},
		&cli.StringFlag{
			Name:    CmdFlag_OrderType,
			Aliases: []string{"t"},
			Value:   string(types.OrderTypeMarket),
		},
		&cli.StringFlag{
			Name:    CmdFlag_TradeMode,
			Aliases: []string{"m"},
			Value:   string(types.TradeModeIsolated),
		},
		&cli.StringFlag{
			Name:    CmdFlag_PosSideType,
			Aliases: []string{"T"},
			Value:   "", //不要设置默认值(仅必须的情况下设置)
		},
	},
	Action: func(ctx *cli.Context) error {
		var err error
		var c *config.Config
		c, err = loadConfig(ctx)
		if err != nil {
			log.Panic(err.Error())
		}
		cexName := types.CexName(ctx.String(CmdFlag_Cex))
		symbol := ctx.String(CmdFlag_Symbol)
		sz := sqlca.NewDecimal(ctx.String(CmdFlag_Sz))
		lever := ctx.String(CmdFlag_Lever)
		orderType := types.OrderType(ctx.String(CmdFlag_OrderType))
		posSideType := types.PositionSideType(ctx.String(CmdFlag_PosSideType))
		tradeMode := types.TradeMode(ctx.String(CmdFlag_TradeMode))

		var px sqlca.Decimal
		if ctx.IsSet(CmdFlag_Px) {
			px = sqlca.NewDecimal(ctx.String(CmdFlag_Px))
		}
		var opts []options.TradeOption

		if lever != "" && lever != "0" {
			opts = append(opts, options.WithLever(lever))
		}
		if orderType != "" {
			opts = append(opts, options.WithOrderType(orderType))
		}
		if posSideType != "" {
			opts = append(opts, options.WithPositionSide(posSideType))
		}
		if tradeMode != "" {
			opts = append(opts, options.WithTradeMode(tradeMode))
		}
		if px.GreaterThan(0) {
			opts = append(opts, options.WithPrice(px.String()))
		}

		cex := api.NewCex(cexName, c)
		var orders []*types.OrderDetail

		orders, err = cex.OpenPosition(context.Background(), symbol, sz, opts...)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("open position", orders)
		return nil
	},
}

var cmdPosClose = &cli.Command{
	Name:    CmdName_Close,
	Aliases: []string{"c"},
	Usage:   "close position",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Symbol,
			Usage:   "example PEPE-USDT",
			Aliases: []string{"s"},
			Value:   types.PEPEUSDT,
		},
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
		&cli.StringFlag{
			Name:  CmdFlag_OrderId,
			Usage: "client order id",
		},
	},
	Action: func(ctx *cli.Context) error {
		var err error
		var c *config.Config
		c, err = loadConfig(ctx)
		if err != nil {
			log.Panic(err.Error())
		}
		strCexName := types.CexName(ctx.String(CmdFlag_Cex))
		cex := api.NewCex(strCexName, c)
		symbol := ctx.String(CmdFlag_Symbol)
		orderId := ctx.String(CmdFlag_OrderId)
		var opts []options.TradeOption

		if orderId != "" {
			opts = append(opts, options.WithCliOrdId(orderId))
		}
		var orders []*types.ClosePositionDetail
		orders, err = cex.ClosePosition(context.Background(), symbol, opts...)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("close position", orders)
		return nil
	},
}

var cmdPosList = &cli.Command{
	Name:    CmdName_List,
	Aliases: []string{"l"},
	Usage:   "position list",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
	},
	Action: func(ctx *cli.Context) error {
		var err error
		var c *config.Config
		c, err = loadConfig(ctx)
		if err != nil {
			log.Panic(err.Error())
		}
		strCexName := types.CexName(ctx.String(CmdFlag_Cex))
		cex := api.NewCex(strCexName, c)
		var orders []*types.OrderListDetail
		orders, err = cex.GetPosition(context.Background())
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("position list", orders)
		return nil
	},
}
