package auth

func (ac *AuthController) initRoutes() {
	authGroup := ac.router.Group("/auth")
	{
		authGroup.POST("/login", ac.middleware.JWT, ac.LoginController)
		// 추가
	}
}
