package controllers

import (
	"alta-store/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func CreateCustomerController(c echo.Context) error {
	customer, err := database.CreateCustomer()
	c.Bind(&customer)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new customer",
		"customer": customer,
	})
}
