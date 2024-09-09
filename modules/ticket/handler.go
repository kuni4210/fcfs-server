package ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TicketHandler struct {
	log    *logrus.Logger
	router *gin.Engine
}

func NewTicketHandler(logger *logrus.Logger, router *gin.Engine) *TicketHandler {
	handler := &TicketHandler{
		log:    logger,
		router: router,
	}
	handler.initRoutes()
	return handler
}

func (h *TicketHandler) BookTicketHandler(c *gin.Context) {
	h.log.Infof("book ticket")
	c.JSON(200, gin.H{"message": "success"})
}
