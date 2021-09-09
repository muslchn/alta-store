package route

import (
	"alta-store/constant"
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
	e.GET("/products", controller.GetProductsByCategoryController)

	// Category
	e.GET("/category", controller.GetCategoryController)
	e.GET("/category/:id", controller.GetCategoryByIdController)

	// Customer Authentication
	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)

	// JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	// Customer Auth
	r.GET("/customers/:id", controller.GetCustomerDetailController)

	// Cart
	r.POST("/carts", controller.AddCartItemController)
	r.GET("/carts", controller.GetCartController)
	r.DELETE("/carts/:id", controller.DeleteCartItemController)

	// Checkout
	r.POST("/checkout", controller.CheckoutController)
	r.GET("/checkout", controller.GetCheckoutController)
	r.GET("/checkout/:id", controller.GetCheckoutByIdController)

	// Payment
	r.POST("/payment", controller.PaymentController)
	r.GET("/payment/last", controller.LastPaymentController)
	r.GET("/payment", controller.PaymentHistoryController)
	r.GET("/payment/:id", controller.PaymentByIdController)

	return e
}
