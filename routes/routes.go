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

	return e
}
