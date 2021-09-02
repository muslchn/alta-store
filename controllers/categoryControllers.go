package controllers

import (
	"alta-store/lib/database"
	"net/http"

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
