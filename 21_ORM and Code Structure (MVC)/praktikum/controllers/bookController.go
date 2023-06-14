package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"
	"praktikum/config"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books": books,
	})
}

func GetBookController(c echo.Context) error {
	book, err := database.GetBookById(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id: " + c.Param("id"),
		"book": book,
	})
}

func CreateBookController(c echo.Context) error {
	var book models.Book
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book": book,
	})
}

func DeleteBookController(c echo.Context) error {
	var books []models.Book

	book, _ := database.GetBookById(c.Param("id"))
	if book != nil {
		if err := config.DB.Delete(&books, c.Param("id")).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete book by id: " + c.Param("id"),
		})
	}
	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"message": "record not found",
	})


}

func UpdateBookController(c echo.Context) error {
	var updatedBook models.Book
	c.Bind(&updatedBook)

	book, err := database.GetBookById(c.Param("id"))
	bookData := book.(models.Book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bookData.Title = updatedBook.Title
	bookData.Writer = updatedBook.Writer
	bookData.Publisher = updatedBook.Publisher
	config.DB.Save(&bookData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id: " + c.Param("id"),
		"book": bookData,
	})


}