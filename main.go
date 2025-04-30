package main

import (
	"fmt"
	"github.com/civet148/gocex/internal/cmd"
	"os"
	"os/signal"

	"github.com/civet148/log"
)

const (
	Version     = "0.2.0"
	ProgramName = "gocex"
)

var (
	BuildTime = "2025-04-25"
	GitCommit = "<N/A>"
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

func main() {
	grace()
	cmd.AppStart(ProgramName, Version, BuildTime, GitCommit)
}
