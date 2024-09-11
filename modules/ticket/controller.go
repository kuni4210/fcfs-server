package ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TicketController struct {
	log           *logrus.Logger
	router        *gin.Engine
	ticketService *TicketService
}

func NewTicketController(logger *logrus.Logger, router *gin.Engine, ticketService *TicketService) *TicketController {
	controller := &TicketController{
		log:           logger,
		router:        router,
		ticketService: ticketService,
	}
	controller.initRoutes()
	return controller
}

func (tc *TicketController) BookTicketController(c *gin.Context) {
	tc.log.Infof("book ticket")
	c.JSON(200, gin.H{"message": "success"})
}
