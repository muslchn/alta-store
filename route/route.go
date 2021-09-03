package route

import (
	constants "alta-store/constant"
	controllers "alta-store/controller"

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

	// Category
	e.GET("/category", controllers.GetCategoryController)
	e.GET("/category/:id", controllers.GetCategoryByIdController)

	// Customer Authentication
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginCustomersController)

	// JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Customer Auth
	r.GET("/customers/:id", controllers.GetCustomerDetailController)

	// Cart
	r.POST("/cart/:id", controllers.AddToCartController)
	r.GET("/cart", controllers.GetCartController)

	// Checkout

	// Payment

	return e
}
