package ticket

func (h *TicketController) initRoutes() {
	ticketGroup := h.router.Group("/ticket")
	{
		ticketGroup.POST("/book", h.BookTicketController)
		// 추가
	}
}
