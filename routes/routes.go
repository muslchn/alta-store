package routes

import (
	"alta-store/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/products", controllers.GetProductsController)
	// e.POST("/customers", controllers.RegisterController)
	e.GET("/products/:id", controllers.GetProductsByCategory)
	e.GET("/cart", controllers.GetCartController)
	e.POST("/cart/:id", controllers.AddToCart)

	return e
}
