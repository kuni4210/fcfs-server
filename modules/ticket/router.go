package ticket

func (tc *TicketController) initRoutes() {
	ticketGroup := tc.router.Group("/ticket")
	{
		ticketGroup.POST("/book", tc.BookTicketController)
		// 추가
	}
}
