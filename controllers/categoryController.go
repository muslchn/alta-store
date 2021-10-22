package controller

import (
	"alta-store/lib/database"
	"alta-store/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategoryController(c echo.Context) error {
	category, err := database.GetCategory()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get categories",
		Data:    category,
	})
}

func GetCategoryByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := database.GetCategoryById(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if category == nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "fail",
			Message: "category not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get category",
		Data:    category,
	})
}
