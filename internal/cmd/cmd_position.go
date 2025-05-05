package cmd

import (
	"context"
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"github.com/urfave/cli/v2"
)

var cmdPos = &cli.Command{
	Name:    CmdName_Pos,
	Aliases: []string{""},
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
			Name:    CmdFlag_InstId,
			Aliases: []string{"s"},
			Value:   types.PEPEUSDT,
		},
		&cli.StringFlag{
			Name:    CmdFlag_Px,
			Aliases: []string{"p"}, //目标代币价格(市价单不需要)
		},
		&cli.StringFlag{
			Name:     CmdFlag_Sz, //购买的数量(当指定
			Aliases:  []string{"z"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    CmdFlag_Lever, //杠杆倍数（以app为准，需要设置账户杠杆信息）
			Aliases: []string{"l"},
			Value:   "1",
		},
		&cli.StringFlag{
			Name:    CmdFlag_OrderType,
			Aliases: []string{"t"},
			Value:   string(types.OrderTypeMarket), //市价/限价...
		},
		&cli.StringFlag{
			Name:    CmdFlag_TradeMode,
			Aliases: []string{"m"},
			Value:   string(types.TradeModeIsolated), //逐仓/全仓
		},
		&cli.StringFlag{
			Name:    CmdFlag_PosSideType,
			Aliases: []string{"T"},
			Value:   "", //不要设置默认值(仅必须的情况下设置)
		},
		&cli.StringFlag{
			Name:    CmdFlag_SideType,
			Aliases: []string{"S"},
			Value:   string(types.SideTypeBuy), //buy=买多 sell=买空
		},
		&cli.StringFlag{
			Name:    CmdFlag_TargetCcy, //市价单交易货币,将sz视为USD数量（限价单sz依然代表合约张数）
			Aliases: []string{"C"},
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
		instId := ctx.String(CmdFlag_InstId)
		sz := sqlca.NewDecimal(ctx.String(CmdFlag_Sz))
		lever := ctx.String(CmdFlag_Lever)
		orderType := types.OrderType(ctx.String(CmdFlag_OrderType))
		posSideType := types.PositionSideType(ctx.String(CmdFlag_PosSideType))
		tradeMode := types.TradeMode(ctx.String(CmdFlag_TradeMode))
		sideType := types.SideType(ctx.String(CmdFlag_SideType))
		tgtCcy := ctx.String(CmdFlag_TargetCcy)

		var px sqlca.Decimal
		if ctx.IsSet(CmdFlag_Px) {
			px = sqlca.NewDecimal(ctx.String(CmdFlag_Px))
		}
		var opts []options.TradeOption

		if lever != "" && lever != "0" {
			opts = append(opts, options.WithLeverage(lever))
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
		if tgtCcy != "" {
			opts = append(opts, options.WithTargetCcy(tgtCcy)) //目标货币（例如：USDT）
		}
		cex := api.NewCex(cexName, c)
		var orders []*types.OrderDetail

		orders, err = cex.OpenPosition(context.Background(), instId, sideType, sz, opts...)
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
			Name:    CmdFlag_InstId,
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
		instId := ctx.String(CmdFlag_InstId)
		orderId := ctx.String(CmdFlag_OrderId)
		var opts []options.TradeOption

		if orderId != "" {
			opts = append(opts, options.WithCliOrdId(orderId))
		}
		var orders []*types.ClosePositionDetail
		orders, err = cex.ClosePosition(context.Background(), instId, opts...)
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
		&cli.StringFlag{
			Name:    CmdFlag_InstId,
			Usage:   "example PEPE-USDT",
			Aliases: []string{"s"},
			Value:   types.PEPEUSDT,
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
		instId := ctx.String(CmdFlag_InstId)
		var orders []*types.OrderListDetail
		orders, err = cex.GetPosition(context.Background(), instId)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("position list", orders)
		return nil
	},
}
