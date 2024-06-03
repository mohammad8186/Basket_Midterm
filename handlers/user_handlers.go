package handlers

import (
	"net/http"

	"github.com/09184417478m/Basket_Midterm/database"
	"github.com/09184417478m/Basket_Midterm/models"
	"github.com/09184417478m/Basket_Midterm/token"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUserIDFromContext(c echo.Context) uint {
	userToken := c.Get("user").(*jwt.Token)
	claims, _ := token.ParseToken(userToken.Raw)
	return claims.UserID
}
func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusCreated, user)
}
func Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	tokenString, err := token.GenerateToken(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
