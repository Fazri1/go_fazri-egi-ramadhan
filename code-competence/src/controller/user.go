package controller

import (
	"net/http"
	"toko-elektronik/lib/database"
	"toko-elektronik/middleware"
	"toko-elektronik/model"

	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) RegisterController(c echo.Context) error {
	var user model.User
	c.Bind(&user)

	id := database.CheckEmail(user.Email)
	if id != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email already registered",
		})
	}

	err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	userResponse := model.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
		"user":    userResponse,
	})
}

func (ctrl *Controller) LoginController(c echo.Context) error {
	var user model.User
	c.Bind(&user)

	err := database.Login(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed login",
		})
	}

	userResponse := model.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
		"user":    userResponse,
		"token":   token,
	})
}
