package controller

import (
	"alta-store/lib/database"
	"alta-store/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCartController(c echo.Context) error {
	products, err := database.GetCart()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}

func AddCartItemController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	payloadData := make(map[string]string)
	payloadData["product_id"] = c.FormValue("product_id")
	payloadData["qty"] = c.FormValue("qty")
	cartItem, err := database.AddCartItem(payloadData, uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"item":   cartItem,
	})
}
