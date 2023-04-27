package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful_api_testing/config"
	"restful_api_testing/constants"
	"restful_api_testing/models"

	"testing"

	jwtMid "github.com/labstack/echo-jwt"
	"github.com/stretchr/testify/assert"
)

func InsertDataBookForTesting() error {
	book := models.Book{
		Title:     "Book Testing",
		Writer:    "egi",
		Publisher: "gramedoa",
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil

}

func TestCreateBookController(t *testing.T) {
	var testCase = Case{
		name:       "Test create book",
		path:       "/jwt/books",
		expectCode: http.StatusOK,
		sizeData:   1,
	}

	payload := models.Book{
		Title:     "Book Testing",
		Writer:    "egi",
		Publisher: "gramedoa",
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}
	eJWT.POST("/books", CreateBookController)
	req := httptest.NewRequest(http.MethodPost, testCase.path, bytes.NewBuffer(buf))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	e.ServeHTTP(rec, req)
	assert.Equal(t, testCase.expectCode, rec.Code)
}

func TestGetBooksController(t *testing.T) {
	var testCase = Case{
		name:       "get books normal",
		path:       "/jwt/books",
		expectCode: http.StatusOK,
		sizeData:   1,
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBookForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.GET("/books", GetBooksController)
	req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+token)

	e.ServeHTTP(rec, req)

	assert.Equal(t, testCase.expectCode, rec.Code)
}

func TestGetBookController(t *testing.T) {
	testCases := []Case{
		{
			name:       "success get book by id",
			path:       "/jwt/books/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed get book by id",
			path:       "/jwt/books/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBookForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.GET("/books/:id", GetBookController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodGet, testCases[i].path, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)

	}
}

func TestDeleteBookController(t *testing.T) {
	testCases := []Case{
		{
			name:       "Success delete book",
			path:       "/jwt/books/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed delete book",
			path:       "/jwt/books/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBookForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.DELETE("/books/:id", DeleteBookController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodDelete, testCases[i].path, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)
	}
}

func TestUpdateBookController(t *testing.T) {
	testCases := []Case{
		{
			name:       "success update book",
			path:       "/jwt/books/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed update book",
			path:       "/jwt/books/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	payload := struct {
		Title     string `json:"title"`
		Writer    string `json:"writer"`
		Publisher string `json:"publisher"`
	}{
		Title:     "Testing Book",
		Writer:    "Fazri Egi",
		Publisher: "Gramedoa",
	}
	buf, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err.Error())
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	InsertDataBookForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.PUT("/books/:id", UpdateBookController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodPut, testCases[i].path, bytes.NewBuffer(buf))
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)
	}
}
