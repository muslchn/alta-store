package controller

import (
	"alta-store/lib/database"
	"alta-store/middleware"
	"alta-store/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCartItemController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	item := make(map[string]string)
	item["productId"] = c.FormValue("productId")
	item["qty"] = c.FormValue("qty")
	validation, cartItem, err := database.AddCartItem(item, uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !validation[0] {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "product not found",
		})
	} else if !validation[1] {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "product out of stock",
		})
	} else if !validation[2] {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "product is less than you want to take",
		})
	}

	return c.JSON(http.StatusCreated, models.Response{
		Status:  "ok",
		Message: "success add item to cart",
		Data:    cartItem,
	})
}

func GetCartController(c echo.Context) error {
	customerId := middleware.ExtractTokenCustomerId(c)
	cart, err := database.GetCart(uint(customerId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "ok",
		Message: "success get cart",
		Data:    cart,
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
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "item not found",
		})
	}

	database.ReturnStock(validateItem["productId"], validateItem["qty"])
	database.DeleteCartItem(uint(cartItemId))

	return c.JSON(http.StatusOK, models.Response{
		Status:  "ok",
		Message: "item removed",
	})
}
