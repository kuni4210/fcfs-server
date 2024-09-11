package app

import (
	"context"
	"database/sql"
	"fcfs-server/config"
	"fcfs-server/middlewares"
	"fcfs-server/modules/auth"
	"fcfs-server/modules/ticket"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	router     *gin.Engine
	httpServer *http.Server
	postgres   *sql.DB
}

func NewApp(log *logrus.Logger, cfg *config.Config) (*App, error) {
	// postgres db 연결
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName)
	postgres, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := postgres.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// router 초기화
	router := gin.Default()

	// return
	return &App{
		log:    log,
		cfg:    cfg,
		router: router,
		httpServer: &http.Server{
			Addr:    ":" + cfg.Server.Port,
			Handler: router,
		},
		postgres: postgres,
	}, nil
}

func (a *App) Run() error {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	a.log.Infof("Starting server on %s", "8080")

	// Middleware 초기화
	middleware := middlewares.NewMiddleware(a.postgres, a.cfg)
	// a.router.Use(middleware.JWT)

	// Service 초기화
	authService := auth.NewAuthService(a.postgres, a.cfg)
	ticketService := ticket.NewTicketService(a.postgres, a.cfg)

	// Controller 초기화
	_ = auth.NewAuthController(a.log, a.router, authService, middleware)
	_ = ticket.NewTicketController(a.log, a.router, ticketService, middleware)

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
