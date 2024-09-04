package auth

import "github.com/gin-gonic/gin"

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	h.log.Infof("login")
	c.JSON(200, gin.H{"message": "success"})
}
