package ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TicketController struct {
	log    *logrus.Logger
	router *gin.Engine
}

func NewTicketController(logger *logrus.Logger, router *gin.Engine) *TicketController {
	handler := &TicketController{
		log:    logger,
		router: router,
	}
	handler.initRoutes()
	return handler
}

func (h *TicketController) BookTicketController(c *gin.Context) {
	h.log.Infof("book ticket")
	c.JSON(200, gin.H{"message": "success"})
}
