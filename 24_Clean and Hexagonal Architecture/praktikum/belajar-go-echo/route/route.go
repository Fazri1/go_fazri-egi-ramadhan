package route

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)

	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	e.GET("/users", controller.GetAllUsers(db))
	e.POST("/users", userController.Create)
}
