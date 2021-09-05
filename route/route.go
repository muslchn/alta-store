package route

import (
	constants "alta-store/constant"
	"alta-store/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Welcome
	e.GET("", controller.WelcomeController)
	e.GET("/welcome", controller.WelcomeController)

	// Products
	e.GET("/products", controller.GetProductsController)
	e.GET("/products/:categoryId", controller.GetProductsByCategoryController)

	// Category
	e.GET("/category", controller.GetCategoryController)
	e.GET("/category/:id", controller.GetCategoryByIdController)

	// Customer Authentication
	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginCustomersController)

	// JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Customer Auth
	r.GET("/customers/:id", controller.GetCustomerDetailController)

	// Cart
	r.POST("/carts", controller.AddCartItemController)
	r.GET("/cart", controller.GetCartController)

	// Checkout

	// Payment

	return e
}
