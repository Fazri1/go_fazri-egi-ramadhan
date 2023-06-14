package database

import (
	"restful_api_testing/config"
	"restful_api_testing/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func GetBookById(id interface{}) (interface{}, error) {
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}
