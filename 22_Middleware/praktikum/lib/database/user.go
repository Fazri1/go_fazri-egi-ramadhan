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

func Login(email, password string) error {
	var user models.User
	var err error
	if err = config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return err
	}

	return nil
}