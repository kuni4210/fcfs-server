package ticket

func (h *TicketHandler) initRoutes() {
	ticketGroup := h.router.Group("/ticket")
	{
		ticketGroup.POST("/book", h.BookTicketHandler)
		// 추가
	}
}
