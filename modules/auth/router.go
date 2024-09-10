package auth

func (h *AuthController) initRoutes() {
	authGroup := h.router.Group("/auth")
	{
		authGroup.POST("/login", h.LoginController)
		// 추가
	}
}
