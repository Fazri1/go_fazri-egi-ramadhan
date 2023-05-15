package database

import (
	"toko-elektronik/config"
	"toko-elektronik/model"
)

func SaveCategory(category *model.Category) error {
	if err := config.DB.Save(&category).Error; err != nil {
		return err
	}

	return nil
}
