package controller

import (
	"alta-store/lib/database"
	"alta-store/middleware"
	"alta-store/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PaymentController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	checkoutId, _ := strconv.Atoi(c.FormValue("checkoutId"))
	payment, err := database.Payment(uint(customerId), uint(checkoutId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if payment == nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "fail",
			Message: "checkout data not found",
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		Status:  "ok",
		Message: "payment success",
		Data:    payment,
	})
}

func LastPaymentController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	lastPayment, err := database.LastPayment(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get last payment",
		Data:    lastPayment,
	})
}

func PaymentHistoryController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	paymentHistory, err := database.PaymentHistory(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get payment history(s)",
		Data:    paymentHistory,
	})
}

func PaymentByIdController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	id, _ := strconv.Atoi(c.Param("id"))
	paymentById, err := database.PaymentById(uint(id), uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if paymentById == nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "fail",
			Message: "data not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get payment history",
		Data:    paymentById,
	})
}
