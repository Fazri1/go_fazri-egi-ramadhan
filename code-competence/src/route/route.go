package route

import (
	"toko-elektronik/constant"
	"toko-elektronik/controller"
	mid "toko-elektronik/middleware"

	jwtMid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	var reqController controller.Controller
	e := echo.New()
	mid.LogMiddleware(e)
	e.POST("/register", reqController.RegisterController)
	e.POST("/login", reqController.LoginController)

	eJWT := e.Group("/auth")
	eJWT.Use(jwtMid.JWT([]byte(constant.SECRET_JWT)))
	eJWT.POST("/items", reqController.AddItemController)
	eJWT.GET("/items", reqController.GetItemsController)
	eJWT.GET("/items/:id", reqController.GetItemByIDController)
	eJWT.PUT("/items/:id", reqController.UpdateItemController)
	eJWT.DELETE("/items/:id", reqController.DeleteItemController)
	eJWT.GET("/items/category/:category_id", reqController.GetItemsByCategoryController)
	eJWT.GET("/items", reqController.SearchItemController)
	eJWT.POST("/category", reqController.AddCategoryController)

	return e
}
