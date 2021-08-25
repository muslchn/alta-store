package controllers

import (
	"alta-store/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetProductsController(c echo.Context) error {
	products, err := database.GetProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}
