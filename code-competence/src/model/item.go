package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Code        string   `json:"Code" form:"Code" gorm:"unique;type:varchar(50)"`
	Name        string   `json:"Name" form:"Name" gorm:"type:varchar(50)"`
	Description string   `json:"Description" form:"Description" gorm:"type:text"`
	CategoryID  uint     `json:"CategoryID" form:"CategoryID"`
	Stock       uint     `json:"Stock" form:"Stock"`
	Price       float64  `json:"Price" form:"Price" gorm:"type:double"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}

type DetailItemResponse struct {
	ID          uint    `json:"ID" form:"ID"`
	Code        string  `json:"Code" form:"Code"`
	Name        string  `json:"Name" form:"Name"`
	Description string  `json:"Description" form:"Description"`
	CategoryID  uint    `json:"CategoryID" form:"CategoryID"`
	Stock       uint    `json:"Stock" form:"Stock"`
	Price       float64 `json:"Price" form:"Price"`
	Category    CategoryResponse
}

type ItemResponse struct {
	ID          uint    `json:"ID" form:"ID"`
	Code        string  `json:"Code" form:"Code"`
	Name        string  `json:"Name" form:"Name"`
	Description string  `json:"Description" form:"Description"`
	CategoryID  uint    `json:"CategoryID" form:"CategoryID"`
	Stock       uint    `json:"Stock" form:"Stock"`
	Price       float64 `json:"Price" form:"Price"`
}

type Category struct {
	gorm.Model
	Name string `json:"Name" form:"Name"`
}

type CategoryResponse struct {
	ID   uint
	Name string
}
