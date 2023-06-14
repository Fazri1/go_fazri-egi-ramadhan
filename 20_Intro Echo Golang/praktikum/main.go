package main

import (
	"praktikum/handlers"

	"github.com/labstack/echo/v4"
)

// ---------------------------------------------------
func main() {
	e := echo.New()

	getUsers := handlers.GetUsersHandler
	getUser := handlers.GetUserHandler
	deleteUser := handlers.DeleteUserHandler
	updateUser := handlers.UpdateUserHandler
	createUser := handlers.CreateUserHandler

	// routing with query parameter
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.DELETE("/users/:id", deleteUser)
	e.PUT("/users/:id", updateUser)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}