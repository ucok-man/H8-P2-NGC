package api

func (app *Application) Routes() *Application {
	users := app.echo.Group("/users")
	{
		users.POST("/register", app.registerHandler)
		users.POST("/login", app.loginHandler)
	}

	return app
}
