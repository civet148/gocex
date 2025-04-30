package cmd

import (
	"context"
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
)

var cmdPub = &cli.Command{
	Name:    CmdName_Public,
	Aliases: []string{""},
	Usage:   "public command",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CmdFlag_Cex,
			Usage:   "",
			Aliases: []string{"x"},
			Value:   string(types.CexNameOkex),
		},
	},
	Subcommands: []*cli.Command{
		cmdInst,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}

var cmdInst = &cli.Command{
	Name:    CmdName_Inst,
	Aliases: []string{""},
	Usage:   "instrument command",
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
		&cli.StringFlag{
			Name:    CmdFlag_InstType,
			Aliases: []string{"t"},
			Usage:   "SPOT/SWAP/FUTURE/OPTION",
			Value:   string(types.InstType_SWAP),
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
		instType := types.InstType(ctx.String(CmdFlag_InstType))
		cex := api.NewCex(cexName, c)
		insts, err := cex.GetInstrument(context.Background(), instId, instType)
		log.Json("instruments", insts)
		return nil
	},
}
