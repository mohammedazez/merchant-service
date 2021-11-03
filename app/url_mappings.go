package app

func mapUrls() {
	api := router.Group("api")
	{
		usersPath := api.Group("/users")
		{
			// users
			usersPath.POST("/register")
			usersPath.POST("/login")
			usersPath.GET("/list")
			usersPath.GET("/list/:user_id")
			usersPath.DELETE("/list/:id")
			usersPath.PUT("/list/:id")
		}

	}
}
