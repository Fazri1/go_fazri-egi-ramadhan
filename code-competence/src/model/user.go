package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"Name" form:"Name" gorm:"type:varchar(255)"`
	Email    string `json:"Email" form:"Email" gorm:"unique"`
	Password string `json:"Password" form:"Password"`
}

type UserResponse struct {
	ID    uint
	Name  string
	Email string
}
