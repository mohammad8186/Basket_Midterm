package handlers

import (
	"net/http"

	"github.com/09184417478m/Basket_Midterm/database"
	"github.com/09184417478m/Basket_Midterm/models"
	"github.com/labstack/echo/v4"
)

func GetBaskets(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	var baskets []models.Basket
	result := database.DB.Where("user_id = ?", userID).Find(&baskets)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, baskets)
}

func CreateBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	var basket models.Basket
	if err := c.Bind(&basket); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	basket.UserID = userID
	result := database.DB.Create(&basket)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusCreated, basket)
}

func UpdateBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	id := c.Param("id")
	var basket models.Basket
	if err := c.Bind(&basket); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result := database.DB.Model(&models.Basket{ID: id, UserID: userID}).Updates(basket)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, basket)
}

func GetBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	id := c.Param("id")
	var basket models.Basket
	result := database.DB.First(&basket, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, basket)
}

func DeleteBasket(c echo.Context) error {
	userID := GetUserIDFromContext(c)
	id := c.Param("id")
	result := database.DB.Delete(&models.Basket{}, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
