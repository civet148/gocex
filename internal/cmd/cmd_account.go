package cmd

import (
	"context"
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
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
