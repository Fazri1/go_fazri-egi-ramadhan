package handlers

import (
	"net/http"
	"praktikum/controllers"
	"praktikum/users"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersHandler(c echo.Context) error {
	allUsers := controllers.GetUsersController()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    allUsers,
	})
	
}

// get user by id
func GetUserHandler(c echo.Context) error {
	idUser, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	message, user := controllers.GetUserController(idUser)

	if user == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": message,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": message,
		"user":     user,
	})
}

// delete user by id
func DeleteUserHandler(c echo.Context) error {
	idUser, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	message, user := controllers.DeleteUserController(idUser)

	if user == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": message,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": message,
		"user":     user,
	})
}

// update user by id
func UpdateUserHandler(c echo.Context) error {
	idUser, err := strconv.Atoi(c.Param("id"))
	updatedUser := users.User{Id: idUser}
	c.Bind(&updatedUser)

	if err != nil {
		return err
	}

	message, user := controllers.UpdateUserController(idUser, updatedUser)

	if user == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": message,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": message,
		"user":     user,
	})
}

// create new user
func CreateUserHandler(c echo.Context) error {
	user := users.User{}
	c.Bind(&user)

	message, createdUser := controllers.CreateUserController(user)

	if message != "success create user" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": message,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
		"user":    createdUser,
	})
}