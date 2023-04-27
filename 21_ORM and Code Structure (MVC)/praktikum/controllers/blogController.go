package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"
	"praktikum/config"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	blogs, err := database.GetBlogs()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs": blogs,
	})
}

func GetBlogController(c echo.Context) error {
	blog, err := database.GetBlogById(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id: " + c.Param("id"),
		"blog": blog,
	})
}

func CreateBlogController(c echo.Context) error {
	var blog models.Blog
	var createdBlog models.BlogResponse
	c.Bind(&createdBlog)

	_, err := database.GetUserById(createdBlog.IDUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "user doesn't exist!",
		})
	}
	blog.IDUser = createdBlog.IDUser
	blog.Title = createdBlog.Title
	blog.Content = createdBlog.Content
	
	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog": createdBlog,
	})
}

func DeleteBlogController(c echo.Context) error {
	var blogs []models.Blog
	blog, _ := database.GetBlogById(c.Param("id"))

	if blog != nil {
		if err := config.DB.Delete(&blogs, c.Param("id")).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete blog by id: " + c.Param("id"),
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "record not found",
	})
}

func UpdateBlogController(c echo.Context) error {
	var blog models.Blog
	updatedBlog := models.BlogResponse{}
	c.Bind(&updatedBlog)

	if err := config.DB.Find(&blog, c.Param("id")).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := database.GetUserById(updatedBlog.IDUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "user doesn't exist!",
		})
	}
	
	blog.IDUser = updatedBlog.IDUser
	blog.Title = updatedBlog.Title
	blog.Content = updatedBlog.Content
	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog by id: " + c.Param("id"),
		"blog": updatedBlog,
	})
}