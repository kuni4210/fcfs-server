package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	log    *logrus.Logger
	router *gin.Engine
}

func NewAuthController(logger *logrus.Logger, router *gin.Engine) *AuthController {
	handler := &AuthController{
		log:    logger,
		router: router,
	}
	handler.initRoutes()
	return handler
}

func (h *AuthController) LoginController(c *gin.Context) {
	h.log.Infof("login")
	c.JSON(200, gin.H{"message": "success"})
}
