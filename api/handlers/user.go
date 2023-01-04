package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Watson-Sei/face-to-face/database"
	"github.com/Watson-Sei/face-to-face/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetUsers returns all users
func GetUsers(c echo.Context) error {
	var users []models.User
	if err := database.DB.Db.Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(200, users)
}

// GetUser returns a user
func GetUser(c echo.Context) error {
	var user models.User
	if err := database.DB.Db.Where("id = ? AND deleted_at IS NULL", c.Param("id")).First(&user).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return c.JSON(http.StatusNotFound, errors.New("User not found"))
		} else {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, errors.New("Internal server error"))
		}
	}
	return c.JSON(200, user)
}

// CreateUser creates a user
func CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Db.Create(&user)
	return c.JSON(200, user)
}

// DeleteUser deletes a user
func DeleteUser(c echo.Context) error {
	database.DB.Db.Delete(&models.User{}, c.Param("id"))
	return c.NoContent(http.StatusNoContent)
}
