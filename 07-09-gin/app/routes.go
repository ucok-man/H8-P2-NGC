package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/contract"
)

func (app *Application) routes() *gin.Engine {
	router := gin.Default()

	// config
	{
		router.RedirectTrailingSlash = false
		router.RedirectFixedPath = false
		router.RemoveExtraSlash = false
		router.NoRoute(func(ctx *gin.Context) {
			contract.ErrNotFound(ctx)
		})
		router.NoMethod(func(ctx *gin.Context) {
			contract.ErrMethodNotAllowed(ctx)
		})
	}

	user := router.Group("/users")
	{
		user.POST("/login", app.loginHandler)
		user.POST("/register", app.registerHandler)
	}

	product := router.Group("/products")
	product.Use(app.WithLogin)
	{
		product.GET("", app.getAllProductHandler)
		product.GET("/:id", app.getProductByIdHandler)
		product.POST("", app.createProductHandler)
		product.PUT("/:id", app.updateProductHandler)
		product.DELETE("/:id", app.deleteProductHandler)
	}

	return router
}
