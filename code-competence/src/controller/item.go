package controller

import (
	"net/http"
	"toko-elektronik/lib/database"
	"toko-elektronik/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct{}

func (ctrl *Controller) GetItemsController(c echo.Context) error {
	items, err := database.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	var itemResponses []model.DetailItemResponse
	for i := range items {
		itemResponse := model.DetailItemResponse{ID: items[i].ID, Code: items[i].Code, Name: items[i].Name,
			Description: items[i].Description, CategoryID: items[i].CategoryID, Stock: items[i].Stock,
			Price: items[i].Price, Category: model.CategoryResponse{ID: items[i].Category.ID, Name: items[i].Category.Name}}
		itemResponses = append(itemResponses, itemResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success",
		"Items":   itemResponses,
	})
}

func (ctrl *Controller) AddItemController(c echo.Context) error {
	var item model.Item
	c.Bind(&item)
	item.Code = uuid.NewString()

	err := database.SaveItem(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	itemResponse := model.ItemResponse{ID: item.ID, Code: item.Code, Name: item.Name,
		Description: item.Description, CategoryID: item.CategoryID, Stock: item.Stock,
		Price: item.Price}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success",
		"Item":    itemResponse,
	})
}

func (ctrl *Controller) GetItemByIDController(c echo.Context) error {
	item, err := database.GetItemByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	itemResponse := model.DetailItemResponse{ID: item.ID, Code: item.Code, Name: item.Name,
		Description: item.Description, CategoryID: item.CategoryID, Stock: item.Stock,
		Price: item.Price, Category: model.CategoryResponse{ID: item.Category.ID, Name: item.Category.Name}}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success",
		"Item":    itemResponse,
	})
}

func (ctrl *Controller) UpdateItemController(c echo.Context) error {
	var updatedItem model.Item
	c.Bind(&updatedItem)

	item, err := database.GetItemByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	item.Code = updatedItem.Code
	item.Name = updatedItem.Name
	item.Description = updatedItem.Description
	item.CategoryID = updatedItem.CategoryID
	item.Stock = updatedItem.Stock
	item.Price = updatedItem.Price

	err = database.SaveItem(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success",
	})
}

func (ctrl *Controller) DeleteItemController(c echo.Context) error {
	err := database.DeleteItem(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success",
	})
}

func (ctrl *Controller) GetItemsByCategoryController(c echo.Context) error {
	items, err := database.GetItemByCategoryID(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	var itemsResponses []model.DetailItemResponse
	for i := range items {
		itemResponse := model.DetailItemResponse{ID: items[i].ID, Code: items[i].Code, Name: items[i].Name,
			Description: items[i].Description, CategoryID: items[i].CategoryID, Stock: items[i].Stock,
			Price: items[i].Price, Category: model.CategoryResponse{ID: items[i].Category.ID, Name: items[i].Category.Name}}
		itemsResponses = append(itemsResponses, itemResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success",
		"Items":   itemsResponses,
	})
}

func (ctrl *Controller) SearchItemController(c echo.Context) error {
	items, err := database.SearchItem(c.QueryParam("keyword"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}

	var itemsResponses []model.ItemResponse
	for i := range items {
		itemResponse := model.ItemResponse{ID: items[i].ID, Code: items[i].Code, Name: items[i].Name,
			Description: items[i].Description, CategoryID: items[i].CategoryID, Stock: items[i].Stock,
			Price: items[i].Price}
		itemsResponses = append(itemsResponses, itemResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success",
		"Items":   itemsResponses,
	})

}
