package auth

func (h *AuthHandler) initRoutes() {
	authGroup := h.router.Group("/auth")
	{
		authGroup.POST("/login", h.LoginHandler)
		// 추가
	}
}
