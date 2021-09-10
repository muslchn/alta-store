package controller

import (
	"alta-store/lib/database"
	"alta-store/middleware"
	"alta-store/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CheckoutController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	cartId, _ := database.CartCheck(uint(customerId))
	totalItem, totalPrice, err := database.GetTotal(uint(cartId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	checkout, err := database.CreateCheckout(uint(customerId), uint(cartId), totalItem, totalPrice)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	database.ChangeCartStatus(uint(cartId))

	return c.JSON(http.StatusCreated, model.Response{
		Status:  "ok",
		Message: "success checkout",
		Data:    checkout,
	})
}

func GetCheckoutController(c echo.Context) error {
	customerId, _ := strconv.Atoi(c.QueryParam("customer_id"))
	checkout, err := database.GetCheckout(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"checkout": checkout,
	})
}

func GetCheckoutByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	checkout, err := database.GetCheckoutById(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"checkout": checkout,
	})
}
