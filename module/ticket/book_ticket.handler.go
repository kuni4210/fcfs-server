package ticket

import (
	"github.com/gin-gonic/gin"
)

func (h *TicketHandler) BookTicketHandler(c *gin.Context) {
	h.log.Infof("book ticket")
	c.JSON(200, gin.H{"message": "success"})
}
