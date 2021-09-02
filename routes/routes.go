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

	// Products
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:categoryId", controllers.GetProductsByCategoryController)

	// Product Categories

	// Customer Authentication
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginCustomersController)

	// JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Customer Auth
	r.GET("/customers/:id", controllers.GetCustomerDetailControllers)

	// Cart
	r.POST("/carts/:id", controllers.AddToCart)
	r.GET("/carts", controllers.GetCartController)

	// Checkout

	// Payment

	return e
}
