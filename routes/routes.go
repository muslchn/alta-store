package routes

import (
	"alta-store/constants"
	"alta-store/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Welcome
	e.GET("", controllers.WelcomeController)
	e.GET("/welcome", controllers.WelcomeController)

	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:categoryId", controllers.GetProductsByCategory)
	e.GET("/cart", controllers.GetCartController)
	e.POST("/cart/:id", controllers.AddToCart)

	// Customer Authentication
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginCustomersController)

	// JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Customer Auth
	r.GET("/customers/:id", controllers.GetCustomerDetailControllers)

	return e
}
