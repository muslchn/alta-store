package controllers

import (
	"alta-store/lib/database"
	"alta-store/models"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterController(c echo.Context) error {
	// get data form value
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	phone := c.FormValue("phone")
	address := c.FormValue("address")
	city := c.FormValue("city")
	state := c.FormValue("state")
	postalCode := c.FormValue("postalCode")

	var customer models.Customer
	customer.FirstName = firstName
	customer.LastName = lastName
	customer.Username = username
	customer.Email = email
	customer.Password = password
	customer.Phone = phone
	customer.Address = address
	customer.City = city
	customer.State = state
	customer.PostalCode = postalCode

	register, err := database.Register(customer)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create customer",
		"customer": register,
	})
}

// type Customer struct {
// 	gorm.Model
// 	FirstName  string `gorm:"size:255;not null" json:"firstName" form:"firstName"`
// 	LastName   string `gorm:"size:255" json:"lastName" form:"lastName"`
// 	Username   string `gorm:"size:255;not null;unique" json:"username" form:"username"`
// 	Email      string `gorm:"size:100;not null;unique" json:"email" form:"email"`
// 	Password   string `gorm:"size:100;not null" json:"password" form:"password"`
// 	Phone      string `gorm:"size:15;not null" json:"phone" form:"phone"`
// 	Address    string `gorm:"size:255;not null" json:"address" form:"address"`
// 	City       string `gorm:"size:100;not null" json:"city" form:"city"`
// 	State      string `gorm:"size:100;not null" json:"state" form:"state"`
// 	PostalCode string `gorm:"5;not null" json:"postalCode" form:"postalCode"`
// }
