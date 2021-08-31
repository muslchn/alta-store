package routes

import (
	"alta-store/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/products", controllers.GetProductsController)
	// e.POST("/customers", controllers.RegisterController)
	e.GET("/products/:categoryId", controllers.GetProductsByCategory)
	e.GET("/cart", controllers.GetCartController)
	e.POST("/cart/:id", controllers.AddToCart)
	e.POST("/register", controllers.RegisterController)

	return e
}
