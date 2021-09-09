package controller

import (
	"alta-store/lib/database"
	"alta-store/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	products, err := database.GetProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.Message{
		Status: "success get products",
		Data:   products,
	})
}

func GetProductsByCategoryController(c echo.Context) error {
	categoryId, _ := strconv.Atoi(c.QueryParam("category"))

	products, err := database.GetProductsByCategory(uint(categoryId))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.Message{
		Status: "success get products",
		Data:   products,
	})
}
