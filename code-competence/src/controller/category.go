package controller

import (
	"net/http"
	"toko-elektronik/lib/database"
	"toko-elektronik/model"

	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) AddCategoryController(c echo.Context) error {
	var category model.Category
	c.Bind(&category)

	err := database.SaveCategory(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	categoryResponse := model.CategoryResponse{ID: category.ID, Name: category.Name}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "success",
		"Category": categoryResponse,
	})
}
