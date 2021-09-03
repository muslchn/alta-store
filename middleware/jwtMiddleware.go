package middleware

import (
	constants "alta-store/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(customerId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["customerId"] = customerId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenCustomerId(e echo.Context) int {
	customer := e.Get("customer").(*jwt.Token)
	if customer.Valid {
		claims := customer.Claims.(jwt.MapClaims)
		customerId := int(claims["customerId"].(float64))
		return customerId
	}
	return 0
}
