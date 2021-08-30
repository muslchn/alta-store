package controllers

import (
	"alta-store/models"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterController(c echo.Context) error {
	// get data form value
	firstName := c.FormValue("firstName")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var customer models.Customer
	customer.FirstName = firstName
	customer.Email = email
	customer.Password = password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create customer",
		"customer": customer,
	})
}

// type Customer struct {
// 	gorm.Model
// 	ID        uint   `gorm:"primaryKey"`
// 	FirstName string `json:"firstName" form:"firstName"`
// 	LastName  string `json:"lastName" form:"lastName"`
// 	Username  string `json:"username" form:"username"`
// 	Password  string `json:"password" form:"password"`
// 	Address    string `json:"address" form:"address"`
// 	City       string `json:"city" form:"city"`
// 	State      string `json:"state" form:"state"`
// 	PostalCode uint   `json:"postalCode" form:"postalCode"`
// 	Phone      int    `json:"phone" form:"phone"`
// 	Email      string `json:"email" form:"email"`
// }
