package main

import (
	"fmt"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/logic"
	"github.com/civet148/godotenv"
	"github.com/civet148/log"
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
		Flags:   []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			var c config.Config
			err := godotenv.Load(&c)
			if err != nil {
				log.Errorf("load .env error %s", err)
				return err
			}
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
