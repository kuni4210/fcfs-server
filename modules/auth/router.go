package auth

func (ac *AuthController) initRoutes() {
	authGroup := ac.router.Group("/auth")
	{
		authGroup.POST("/login", ac.LoginController)
		// 추가
	}
}
