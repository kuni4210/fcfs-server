package middlewares

import (
	"database/sql"
	"fcfs-server/config"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	JWT gin.HandlerFunc
}

func NewMiddleware(postgres *sql.DB, cfg *config.Config) *Middleware {
	return &Middleware{
		JWT: JwtMiddleware(postgres, cfg.Jwt.SecretKey),
	}
}
