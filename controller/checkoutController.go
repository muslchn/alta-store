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
	customerId := middleware.ExtractTokenCustomerId(c)
	checkout, err := database.GetCheckout(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get checkout(s)",
		Data:    checkout,
	})
}

func GetCheckoutByIdController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	id, _ := strconv.Atoi(c.Param("id"))
	checkout, err := database.GetCheckoutById(uint(id), uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if checkout == nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "fail",
			Message: "checkout data not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get checkout data",
		Data:    checkout,
	})
}
