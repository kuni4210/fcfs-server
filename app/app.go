package app

import (
	"context"
	"fcfs-server/config"
	"fcfs-server/module/auth"
	"fcfs-server/module/ticket"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	router     *gin.Engine
	httpServer *http.Server

	authHandler   *auth.AuthHandler
	ticketHandler *ticket.TicketHandler

	// db connection 추가
	// redisClient *redis.Client
}

func NewApp(log *logrus.Logger, cfg *config.Config) *App {
	router := gin.Default()

	authHandler := auth.NewAuthHandler(log, router)
	ticketHandler := ticket.NewTicketHandler(log, router)

	return &App{
		log:    log,
		cfg:    cfg,
		router: router,
		httpServer: &http.Server{
			Addr:    ":" + cfg.Server.Port,
			Handler: router,
		},

		authHandler:   authHandler,
		ticketHandler: ticketHandler,
	}
}

func (a *App) Run() error {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	a.log.Infof("Starting server on %s", "8080")

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Errorf("HTTP server error: %+v", err)
		}
	}()
	defer a.httpServer.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.log.Infof("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutdownCancel()
	err := a.httpServer.Shutdown(shutdownCtx)
	if err != nil {
		a.log.Errorf("Server shutdown error: %+v", err)
	}
	a.log.Infof("Server exited properly")

	return nil
}
