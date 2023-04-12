package database

import (
	"restful_api_testing/config"
	"restful_api_testing/models"
)

func GetBlogs() (interface{}, error) {
	var blogs []models.Blog
	if err := config.DB.Preload("User").Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func GetBlogById(id interface{}) (interface{}, error) {
	var blog models.Blog
	if err := config.DB.Preload("User").First(&blog, id).Error; err != nil {
		return nil, err
	}

	return blog, nil
}
