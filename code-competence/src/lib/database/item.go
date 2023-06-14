package database

import (
	"toko-elektronik/config"
	"toko-elektronik/model"
)

func GetAllItems() ([]model.Item, error) {
	var items []model.Item
	if err := config.DB.Preload("Category").Find(&items).Error; err != nil {
		return items, err
	}

	return items, nil
}

func SaveItem(item *model.Item) error {
	if err := config.DB.Save(item).Error; err != nil {
		return err
	}

	return nil
}

func GetItemByID(id string) (model.Item, error) {
	var item model.Item
	if err := config.DB.Preload("Category").First(&item, id).Error; err != nil {
		return item, err
	}

	return item, nil
}

func DeleteItem(id string) error {
	var item model.Item
	if err := config.DB.Delete(&item, id).Error; err != nil {
		return err
	}

	return nil
}

func GetItemByCategoryID(id string) ([]model.Item, error) {
	var items []model.Item
	if err := config.DB.Preload("Category").Joins("Items").Where("category_id = ?", id).Find(&items).Error; err != nil {
		return items, err
	}

	return items, nil
}

func SearchItem(name string) ([]model.Item, error) {
	var items []model.Item
	if err := config.DB.Table("items").Where("name LIKE ? AND deleted_at IS NULL", "%"+name+"%").Scan(&items).Error; err != nil {
		return items, err
	}

	return items, nil
}
