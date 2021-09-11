package controller

import (
	"alta-store/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	// Setup
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetProductsController(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/products")

		// Assertions
		if assert.NoError(t, GetProductsController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestGetProductsByCategoryController(t *testing.T) {
	t.Run("test case 1, valid category id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/products")
		c.SetParamNames("category")
		c.SetParamValues("1")

		// Assertions
		if assert.NoError(t, GetProductsByCategoryController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("test case 2, invalid category id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/products")
		c.SetParamNames("category")
		c.SetParamValues("4")

		// Assertions
		if assert.NoError(t, GetProductsByCategoryController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
