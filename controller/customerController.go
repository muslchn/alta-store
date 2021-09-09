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
	var customer model.Customer
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

func LoginController(c echo.Context) error {
	type dataModel struct {
		email string
		token string
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	customer, name, token, err := database.Login(email, password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if customer == nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "fail",
			Message: "account not found or password incorrect",
		})
	}

	hello := "hello" + name

	data := dataModel{
		email: email,
		token: token,
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: hello,
		Data:    data,
	})
}

func GetCustomerDetailController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := database.GetDetailCustomer(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if customer == nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "fail",
			Message: "data not found",
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  "ok",
		Message: "success get customer detail",
		Data:    customer,
	})
}
