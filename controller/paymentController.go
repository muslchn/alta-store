package controller

import (
	"alta-store/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PaymentController(c echo.Context) error {
	checkoutId, _ := strconv.Atoi(c.FormValue("checkoutId"))
	payment, err := database.Payment(uint(checkoutId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"payment": payment,
	})
}

func LastPaymentController(c echo.Context) error {
	customerId, _ := strconv.Atoi(c.QueryParam("customerId"))
	lastPayment, err := database.LastPayment(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"lastPayment": lastPayment,
	})
}

func PaymentHistoryController(c echo.Context) error {
	customerId, _ := strconv.Atoi(c.QueryParam("customerId"))
	paymentHistory, err := database.PaymentHistory(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"paymentHistory": paymentHistory,
	})
}

func PaymentByIdController(c echo.Context) error {
	return nil
}
