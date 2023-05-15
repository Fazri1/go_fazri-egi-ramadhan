package database

import (
	"toko-elektronik/config"
	"toko-elektronik/model"
	"toko-elektronik/util"
)

func CreateUser(user *model.User) error {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = util.ComparePassword(hashedPassword, user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	if err := config.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func Login(user *model.User) error {
	var userForLogin model.User
	if err := config.DB.Table("users").Select("password").Where("email = ?", user.Email).First(&userForLogin).Error; err != nil {
		return err
	}

	err := util.ComparePassword(userForLogin.Password, user.Password)
	if err != nil {
		return err
	}

	if err := config.DB.Where("email = ? AND password = ?", user.Email, userForLogin.Password).First(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckEmail(email string) uint {
	var user model.User
	if err := config.DB.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return 0
	}

	if user.ID != 0 {
		return user.ID
	}

	return 0
}
