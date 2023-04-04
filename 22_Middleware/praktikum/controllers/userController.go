package controllers

import (
	"net/http"
	"praktikum/config"
	"praktikum/lib/database"
	"praktikum/middlewares"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "users":   users,
	})
}
  
// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	user, err := database.GetUserById(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id: " + c.Param("id"),
		"user": user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var users []models.User
	user, _ := database.GetUserById(c.Param("id"))

	if user != nil {
		if err := config.DB.Delete(&users, c.Param("id")).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user by id: " + c.Param("id"),
		})
	} 

	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"message": "record not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user models.User
	updatedUser := models.User{}
	c.Bind(&updatedUser)

	if err := config.DB.Find(&user, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password
	config.DB.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id: " + c.Param("id"),
		"user": user,
	})
}

func LoginController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Username or Password.",
		})
	}

	token, err := middlewares.CreateToken(user.ID, user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "error",
		})
	}

	userResponse := models.UserResponse{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user": userResponse,
	})


}