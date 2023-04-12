package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"restful_api_testing/config"
	"restful_api_testing/constants"
	"restful_api_testing/models"

	"testing"

	jwtMid "github.com/labstack/echo-jwt"
	"github.com/stretchr/testify/assert"
)

func InsertDataBlogForTesting() error {
	blog := models.Blog{
		IDUser:  1,
		Title:   "How to create unit test",
		Content: "Dolar Ipsum",
	}

	if err := config.DB.Save(&blog).Error; err != nil {
		return err
	}
	return nil

}

func TestCreateBlogController(t *testing.T) {
	var testCase = Case{
		name:       "Test create blog",
		path:       "/jwt/blogs",
		expectCode: http.StatusOK,
		sizeData:   1,
	}

	payload := models.Blog{
		IDUser:  1,
		Title:   "Blog Testing",
		Content: "dolar ipsum",
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBlogForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}
	eJWT.POST("/blogs", CreateBlogController)
	req := httptest.NewRequest(http.MethodPost, testCase.path, bytes.NewBuffer(buf))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	e.ServeHTTP(rec, req)
	assert.Equal(t, testCase.expectCode, rec.Code)
}

func TestGetBlogsController(t *testing.T) {
	var testCase = Case{
		name:       "get blogs normal",
		path:       "/jwt/blogs",
		expectCode: http.StatusOK,
		sizeData:   1,
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBlogForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.GET("/blogs", GetBlogsController)
	req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+token)

	e.ServeHTTP(rec, req)

	assert.Equal(t, testCase.expectCode, rec.Code)
}

func TestGetBlogController(t *testing.T) {
	testCases := []Case{
		{
			name:       "success get blog by id",
			path:       "/jwt/blogs/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed get blog by id",
			path:       "/jwt/blogs/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBlogForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.GET("/blogs/:id", GetBlogController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodGet, testCases[i].path, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)

	}
}

func TestDeleteBlogController(t *testing.T) {
	testCases := []Case{
		{
			name:       "Success delete blog",
			path:       "/jwt/blogs/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed delete blog",
			path:       "/jwt/blogs/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBlogForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.DELETE("/blogs/:id", DeleteBlogController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodDelete, testCases[i].path, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)
	}
}

func TestUpdateBlogController(t *testing.T) {
	testCases := []Case{
		{
			name:       "success update blog",
			path:       "/jwt/blogs/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		// {
		// 	name:       "failed update blog",
		// 	path:       "/jwt/blogs/2",
		// 	expectCode: http.StatusBadRequest,
		// 	sizeData:   1,
		// },
	}

	payload := struct {
		IDUser  uint   `json:"iduser"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		IDUser:  1,
		Title:   "Create unit test in golang",
		Content: "dolar ipsum doloris",
	}
	buf, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err.Error())
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBlogForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.PUT("/blogs/:id", UpdateBlogController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodPut, testCases[i].path, bytes.NewBuffer(buf))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		fmt.Println("re", req)

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)
	}
}
