package main

import (
	"fcfs-server/app"
	"fcfs-server/config"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := &logrus.Logger{}
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("could not load config: %v\n", err)
	}

	application := app.NewApp(log, cfg)
	if err := application.Run(); err != nil {
		fmt.Printf("App error: %v\n", err)
		os.Exit(1)
	}
}
