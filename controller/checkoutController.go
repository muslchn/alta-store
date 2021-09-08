package controller

import (
	"alta-store/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CheckoutController(c echo.Context) error {
	cartId, _ := strconv.Atoi(c.FormValue("cartId"))
	totalItem, totalPrice, err := database.GetTotal(uint(cartId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	checkout, err := database.CreateCheckout(uint(cartId), totalItem, totalPrice)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"checkout": checkout,
	})
}

func GetCheckoutController(c echo.Context) error {
	return nil
}

func GetCheckoutByIdController(c echo.Context) error {
	return nil
}
