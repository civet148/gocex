package cmd

import (
	"fmt"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/logic"
	"github.com/civet148/godotenv"
	"github.com/civet148/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	CmdName_Public      = "pub"       //公共参数
	CmdName_Inst        = "inst"      //交易基本参数
	CmdName_Pos         = "pos"       //仓位子命令
	CmdName_Open        = "open"      //建仓
	CmdName_Close       = "close"     //平仓
	CmdName_List        = "list"      //仓位列表
	CmdName_Account     = "acc"       //账户信息
	CmdName_Balance     = "balance"   //账户余额
	CmdName_GetLeverage = "get-lever" //获取账户杠杆信息
	CmdName_SetLeverage = "set-lever" //设置账户杠杆信息
)

const (
	CmdFlag_Config        = "config"
	CmdFlag_Debug         = "debug"
	CmdFlag_Cex           = "cex"
	CmdFlag_InstId        = "inst-id"
	CmdFlag_InstType      = "inst-type"
	CmdFlag_Px            = "px"
	CmdFlag_Sz            = "sz"
	CmdFlag_Lever         = "lever"
	CmdFlag_OrderType     = "order-type"
	CmdFlag_PosSideType   = "pos-side"
	CmdFlag_TradeMode     = "trade-mode"
	CmdFlag_OrderId       = "order-id"
	CmdFlag_SideType      = "side-type"
	CmdFlag_TargetCcy     = "target-ccy"
	CmdFlag_Ccy           = "ccy"
	CmdFlag_Simulate      = "sim"
	CmdFlag_Continuous    = "continuous"
	CmdFlag_RiseThreshold = "rise-threshold"
	CmdFlag_PullbackRate  = "pullback-rate"
	CmdFlag_StopWinPct    = "stop-win-pct"
	CmdFLag_StopLossPct   = "stop-loss-pct"
	CmdFlag_FastRise      = "fast-rise"
	CmdFlag_Leverage      = "leverage"
	CmdFlag_CheckDuration = "check-duration"
)

func AppStart(program, ver, buildTime, commit string) {

	app := &cli.App{
		Name:    program,
		Usage:   "",
		Version: fmt.Sprintf("v%s %s commit %s", ver, buildTime, commit),
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
			&cli.BoolFlag{
				Name:    CmdFlag_Simulate,
				Usage:   "simulate mode",
				Aliases: []string{"s"},
			},
			&cli.IntFlag{
				Name:    CmdFlag_Continuous,
				Usage:   "price rise continuous times",
				Aliases: []string{"t"},
			},
			&cli.Float64Flag{
				Name:    CmdFlag_RiseThreshold,
				Usage:   "price rise threshold rate",
				Aliases: []string{"r"},
			},
			&cli.Float64Flag{
				Name:    CmdFlag_PullbackRate,
				Usage:   "price pull back rate",
				Aliases: []string{"p"},
			},
			&cli.Float64Flag{
				Name:    CmdFlag_StopWinPct,
				Usage:   "stop win percent",
				Aliases: []string{"w"},
			},
			&cli.Float64Flag{
				Name:    CmdFLag_StopLossPct,
				Usage:   "stop loss percent",
				Aliases: []string{"l"},
			},
			&cli.Float64Flag{
				Name:    CmdFlag_FastRise,
				Usage:   "price fast rise rate",
				Aliases: []string{"f"},
			},
			&cli.IntFlag{
				Name:    CmdFlag_Leverage,
				Usage:   "contract leverage",
				Aliases: []string{"L"},
			},
			&cli.DurationFlag{
				Name:    CmdFlag_CheckDuration,
				Usage:   "price check duration",
				Aliases: []string{"D"},
			},
		},
		Commands: []*cli.Command{
			cmdAcc,
			cmdPos,
			cmdPub,
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

func loadConfig(ctx *cli.Context) (*config.Config, error) {
	var c config.Config
	var sim = ctx.Bool(CmdFlag_Simulate)
	var continuous = ctx.Int(CmdFlag_Continuous)
	var riseThreshold = ctx.Float64(CmdFlag_RiseThreshold)
	var pullbackRate = ctx.Float64(CmdFlag_PullbackRate)
	var stopWinPct = ctx.Float64(CmdFlag_StopWinPct)
	var stopLossPct = ctx.Float64(CmdFLag_StopLossPct)
	var fastRise = ctx.Float64(CmdFlag_FastRise)
	var leverage = ctx.Int(CmdFlag_Leverage)
	var checkDuration = ctx.Duration(CmdFlag_CheckDuration)

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
	if sim {
		c.Simulate = true
	}
	if continuous > 0 {
		c.Continuous = int32(continuous)
	}
	if riseThreshold > 0 {
		c.RiseThreshold = riseThreshold
	}
	if pullbackRate > 0 {
		c.PullBackRate = pullbackRate
	}
	if stopWinPct > 0 {
		c.StopWinPct = stopWinPct
	}
	if stopLossPct > 0 {
		c.StopLossPct = stopLossPct
	}
	if fastRise > 0 {
		c.FastRise = fastRise
	}
	if leverage > 0 {
		c.Leverage = int32(leverage)
	}
	if checkDuration > 0 {
		c.CheckDur = checkDuration
	}
	log.Json(&c)
	return &c, nil
}
