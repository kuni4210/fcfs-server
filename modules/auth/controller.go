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

// func (ac *AuthController) LoginController(c *gin.Context) {
// 	ac.log.Infof("login")

// 	var req struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.BindJSON(&req); err != nil {
// 		ac.log.Errorf("error binding request: %v", err)
// 		c.JSON(400, gin.H{"message": "invalid request"})
// 		return
// 	}

// 	token, err := ac.authService.Login(req.Username, req.Password)
// 	if err != nil {
// 		ac.log.Errorf("login error: %v", err)
// 		c.JSON(401, gin.H{"message": "invalid credentials"})
// 		return
// 	}

// 	c.JSON(200, gin.H{"token": token})
// }

func (ac *AuthController) LoginController(c *gin.Context) {
	ac.log.Infof("login")

	c.JSON(200, gin.H{"token": "token"})
}
