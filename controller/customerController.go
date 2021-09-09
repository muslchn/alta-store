package controller

import (
	"alta-store/lib/database"
	"alta-store/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	// // get data form value
	// firstName := c.FormValue("firstName")
	// lastName := c.FormValue("lastName")
	// username := c.FormValue("username")
	// email := c.FormValue("email")
	// password := c.FormValue("password")
	// phone := c.FormValue("phone")
	// address := c.FormValue("address")
	// city := c.FormValue("city")
	// state := c.FormValue("state")
	// postalCode := c.FormValue("postalCode")

	// var customer model.Customer
	// customer.FirstName = firstName
	// customer.LastName = lastName
	// customer.Username = username
	// customer.Email = email
	// customer.Password = password
	// customer.Phone = phone
	// customer.Address = address
	// customer.City = city
	// customer.State = state
	// customer.PostalCode = postalCode

	// binding data
	customer := model.Customer{}
	c.Bind(&customer)
	register, err := database.Register(&customer)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "register succeed",
		Data:    register,
	})
}

func LoginCustomersController(c echo.Context) error {
	customer := model.Customer{}
	c.Bind(&customer)

	customers, err := database.LoginCustomers(&customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success login",
		"customers": customers,
	})
}

func GetCustomerDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	customers, err := database.GetDetailCustomers((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"customers": customers,
	})
}
