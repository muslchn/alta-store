package controller

import (
	"alta-store/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

func GetProductsByCategoryController(c echo.Context) error {
	categoryId, _ := strconv.Atoi(c.Param("categoryId"))

	products, err := database.GetProductsByCategory(categoryId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}
