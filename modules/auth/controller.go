package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	log         *logrus.Logger
	router      *gin.Engine
	authService *AuthService
}

func NewAuthController(logger *logrus.Logger, router *gin.Engine, authService *AuthService) *AuthController {
	controller := &AuthController{
		log:         logger,
		router:      router,
		authService: authService,
	}
	controller.initRoutes()
	return controller
}

func (h *AuthController) LoginController(c *gin.Context) {
	h.log.Infof("login")
	c.JSON(200, gin.H{"message": "success"})
}
