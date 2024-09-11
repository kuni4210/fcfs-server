package main

import (
	"fcfs-server/app"
	"fcfs-server/config"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// config 로드
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("could not load config: %v\n", err)
	}

	// 로그 설정
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(logrus.InfoLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		log.SetLevel(logrus.DebugLevel)
	}

	// App 초기화
	application, err := app.NewApp(log, cfg)
	if err != nil {
		fmt.Printf("App error: %v\n", err)
		os.Exit(1)
	}
	if err = application.Run(); err != nil {
		fmt.Printf("App error: %v\n", err)
		os.Exit(1)
	}
}
