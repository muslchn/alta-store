package controller

import (
	"alta-store/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategoryController(c echo.Context) error {
	category, err := database.GetCategory()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"category": category,
	})
}

func GetCategoryByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := database.GetCategoryById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"category": category,
	})
}
