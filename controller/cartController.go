package controller

import (
	"alta-store/lib/database"
	"alta-store/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

func GetCartController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	cart, err := database.GetCart(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": cart,
	})
}

func DeleteCartItemController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	cartItemId, _ := strconv.Atoi(c.Param("id"))
	validate, validateItem, err := database.ValidateCartItem(uint(cartItemId), uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !validate {
		return c.JSON(http.StatusForbidden, map[string]string{
			"status":  "fail",
			"message": "You do not have access to delete the data",
		})
	}

	database.ReturnStock(validateItem["product_id"], validateItem["qty"])
	database.DeleteCartItem(uint(cartItemId))

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "Product has been removed from the cart",
	})
}
