package controller

import (
	"alta-store/lib/database"
	"net/http"
	"strconv"

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

func AddToCartController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	products, err := database.AddToCart(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}
