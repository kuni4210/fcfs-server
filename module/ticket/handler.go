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

func (h *TicketHandler) initRoutes() {
	ticketGroup := h.router.Group("/ticket")
	{
		ticketGroup.POST("/book", h.BookTicketHandler)
		// 추가
	}
}
