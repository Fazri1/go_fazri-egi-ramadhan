package repository

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(data model.User) error {
	return u.db.Create(&data).Error
}

func (u *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
