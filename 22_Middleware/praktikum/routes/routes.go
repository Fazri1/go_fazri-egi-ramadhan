package routes

import (
	"praktikum/constants"
	"praktikum/controllers"
	mid "praktikum/middlewares"

	"github.com/labstack/echo/v4"
	jwtMid "github.com/labstack/echo-jwt"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	mid.LogMiddleware(e)

	// Route / to handler function
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	// Users endpoint
	eJwt := e.Group("/jwt")
	eJwt.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))
	eJwt.GET("/users", controllers.GetUsersController)
	eJwt.DELETE("/users/:id", controllers.DeleteUserController)
	eJwt.PUT("/users/:id", controllers.UpdateUserController)
	eJwt.GET("/users", controllers.GetUsersController)
	eJwt.GET("/users/:id", controllers.GetUserController)

	// Books endpoint
	eJwt.GET("/books", controllers.GetBooksController)
	eJwt.GET("/books/:id", controllers.GetBookController)
	eJwt.POST("/books", controllers.CreateBookController)
	eJwt.DELETE("/books/:id", controllers.DeleteBookController)
	eJwt.PUT("/books/:id", controllers.UpdateBookController)

	// Blogs endpoint
	eJwt.GET("/blogs", controllers.GetBlogsController)
	eJwt.GET("/blogs/:id", controllers.GetBlogController)
	eJwt.POST("/blogs", controllers.CreateBlogController)
	eJwt.DELETE("/blogs/:id", controllers.DeleteBlogController)
	eJwt.PUT("/blogs/:id", controllers.UpdateBlogController)

	return e
}