package cmd

import (
	"context"
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
)

var cmdAcc = &cli.Command{
	Name:    CmdName_Account,
	Aliases: []string{"a"},
	Usage:   "account command",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
	},
	Subcommands: []*cli.Command{
		cmdBalance,
		cmdGetLever,
		cmdSetLever,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

var cmdBalance = &cli.Command{
	Name:    CmdName_Balance,
	Aliases: []string{"b"},
	Usage:   "balance command",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
		&cli.StringFlag{
			Name:    CmdFlag_Ccy,
			Aliases: []string{"c"},
			Value:   types.USDT,
		},
	},
	Subcommands: []*cli.Command{},
	Action: func(ctx *cli.Context) error {
		var err error
		var c *config.Config
		c, err = loadConfig(ctx)
		if err != nil {
			log.Panic(err.Error())
		}
		cexName := types.CexName(ctx.String(CmdFlag_Cex))
		ccy := ctx.String(CmdFlag_Ccy)
		cex := api.NewCex(cexName, c)
		balance, err := cex.GetBalance(context.Background(), ccy)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("balance", balance)
		return nil
	},
}

var cmdGetLever = &cli.Command{
	Name:    CmdName_GetLeverage,
	Aliases: []string{"gl"},
	Usage:   "get leverage command",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
		&cli.StringFlag{
			Name:    CmdFlag_InstId,
			Aliases: []string{"s"},
			Value:   types.PEPEUSDT,
		},
	},
	Subcommands: []*cli.Command{},
	Action: func(ctx *cli.Context) error {
		var err error
		var c *config.Config
		c, err = loadConfig(ctx)
		if err != nil {
			log.Panic(err.Error())
		}
		cexName := types.CexName(ctx.String(CmdFlag_Cex))
		instId := ctx.String(CmdFlag_InstId)

		var opts []options.TradeOption
		opts = append(opts, options.WithSwap())

		cex := api.NewCex(cexName, c)
		balance, err := cex.GetLeverage(context.Background(), instId, types.MarginModeIsolated, opts...)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("balance", balance)
		return nil
	},
}

var cmdSetLever = &cli.Command{
	Name:    CmdName_SetLeverage,
	Aliases: []string{"sl"},
	Usage:   "set leverage command",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
		&cli.StringFlag{
			Name:    CmdFlag_InstId,
			Aliases: []string{"s"},
			Value:   types.PEPEUSDT,
		},
	},
	Subcommands: []*cli.Command{},
	Action: func(ctx *cli.Context) error {
		var err error
		var c *config.Config
		c, err = loadConfig(ctx)
		if err != nil {
			log.Panic(err.Error())
		}
		cexName := types.CexName(ctx.String(CmdFlag_Cex))
		instId := ctx.String(CmdFlag_InstId)
		lever := ctx.Args().First()
		if lever == "" || lever == "0" {
			return log.Errorf("invalid leverage")
		}
		cex := api.NewCex(cexName, c)
		var opts []options.TradeOption
		opts = append(opts, options.WithLeverage(lever))
		opts = append(opts, options.WithSwap())

		balance, err := cex.SetLeverage(context.Background(), instId, types.MarginModeIsolated, opts...)
		if err != nil {
			return log.Errorf(err.Error())
		}
		log.Json("balance", balance)
		return nil
	},
}
