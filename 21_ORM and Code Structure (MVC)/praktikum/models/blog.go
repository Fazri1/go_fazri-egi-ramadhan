package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	IDUser uint 	`json:"iduser" form:"iduser"`
	Title string	`json:"title" form:"title"`
	Content string	`json:"content" form:"content"`
	User User `gorm:"foreignKey:IDUser"`
}

type BlogResponse struct {
	IDUser uint 	`json:"iduser" form:"iduser"`
	Title string	`json:"title" form:"title"`
	Content string	`json:"content" form:"content"`
}
