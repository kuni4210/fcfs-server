package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	log    *logrus.Logger
	router *gin.Engine
}

func NewAuthHandler(logger *logrus.Logger, router *gin.Engine) *AuthHandler {
	handler := &AuthHandler{
		log:    logger,
		router: router,
	}
	handler.initRoutes()
	return handler
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	h.log.Infof("login")
	c.JSON(200, gin.H{"message": "success"})
}
