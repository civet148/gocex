package main

import (
	"fmt"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/logic"
	"github.com/civet148/godotenv"
	"github.com/civet148/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

const (
	Version     = "0.1.0"
	ProgramName = "gocex"
)

var (
	BuildTime = "2025-04-11"
	GitCommit = "<N/A>"
)

const (
	CmdFlag_Config = "config"
	CmdFlag_Debug  = "debug"
)

func init() {
	log.SetLevel("info")
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
		Action: func(ctx *cli.Context) error {

			if ctx.Bool(CmdFlag_Debug) {
				log.SetLevel("debug")
			}

			var c config.Config

			//设置配置文件
			viper.SetConfigFile(ctx.String(CmdFlag_Config))

			// 读取配置文件
			err := viper.ReadInConfig()
			if err != nil {
				return log.Errorf(err)
			}

			// 反序列化到结构体
			err = viper.Unmarshal(&c)
			if err != nil {
				return log.Errorf(err)
			}
			if c.ApiKey == "" {
				err = godotenv.Load(&c)
				if err != nil {
					log.Errorf("load .env error %s", err)
					return err
				}
			}

			log.Json(&c)
			cex := logic.NewCexLogic(&c)
			return cex.Run()
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}
