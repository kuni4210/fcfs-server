package main

import (
	"fcfs-server/app"
	"fcfs-server/config"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := &config.Config{
		// Server: struct{ Address string }{Address: ":8080"},
	}
	log := &logrus.Logger{}
	app := app.NewApp(log, cfg)

	if err := app.Run(); err != nil {
		log.Errorf("App error: %v", err)
		os.Exit(1)
	}
}
