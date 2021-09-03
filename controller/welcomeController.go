package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func WelcomeController(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to ALTA STORE")
}
