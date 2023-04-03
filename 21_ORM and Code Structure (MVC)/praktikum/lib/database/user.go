package database

import (
	"praktikum/config"
	"praktikum/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	
	return users, nil
	
}

func GetUserById(id interface{}) (interface{}, error) {
	var user models.User

	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}