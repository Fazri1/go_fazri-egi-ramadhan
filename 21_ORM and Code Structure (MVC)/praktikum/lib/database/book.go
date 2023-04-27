package database

import (
	"praktikum/config"
	"praktikum/models"
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

	if err := config.DB.Find(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}