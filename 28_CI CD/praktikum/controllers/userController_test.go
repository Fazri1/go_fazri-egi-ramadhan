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
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	name       string
	path       string
	expectCode int
	sizeData   int
}

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForTesting() error {
	user := models.User{
		Name:     "egi",
		Email:    "egi@gmail.com",
		Password: "egi123",
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil

}

func Login(email, password string, e *echo.Echo) (string, error) {
	var token string
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		email,
		password,
	}
	buf, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	e.POST("/login", LoginController)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(buf))
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		resp := struct {
			Message string
			User    models.UserResponse
		}{}

		err = json.Unmarshal([]byte(rec.Body.String()), &resp)
		if err != nil {
			return "", err
		}
		token = resp.User.Token
	}
	return token, nil
}

func TestCreateUserController(t *testing.T) {
	var testCases = Case{
		name:       "Test create user",
		path:       "/users",
		expectCode: http.StatusOK,
		sizeData:   1,
	}

	payload := models.User{

		Name:     "fazri",
		Email:    "fazri@gmail.com",
		Password: "fazri123",
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	e := InitEchoTestAPI()
	e.POST("/users", CreateUserController)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(buf))
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	e.ServeHTTP(rec, req)
	assert.Equal(t, testCases.expectCode, rec.Code)
}

func TestLogin(t *testing.T) {
	var testCases = []Case{
		{
			name:       "invalid email",
			path:       "/login",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
		{
			name:       "invalid password",
			path:       "/login",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
		{
			name:       "no record",
			path:       "/login",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
		{
			name:       "login success",
			path:       "/login",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	payload := []struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		{
			Email:    "fazri@gmail.com",
			Password: "egi123",
		},
		{
			Email:    "egi@gmail.com",
			Password: "fazri123",
		},
		{
			Email:    "fazri@gmail.com",
			Password: "fazri123",
		},
		{
			Email:    "egi@gmail.com",
			Password: "egi123",
		},
	}

	for i := 0; i < len(testCases); i++ {
		// encode the payload as JSON
		buf, err := json.Marshal(payload[i])
		if err != nil {
			t.Fatal(err)
		}

		e := InitEchoTestAPI()
		InsertDataUserForTesting()
		e.POST("/login", LoginController)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(buf))
		rec := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		e.ServeHTTP(rec, req)

		assert.Equal(t, testCases[i].expectCode, rec.Code)
	}
}

func TestGetUsersController(t *testing.T) {
	var testCases = []Case{
		{
			name:       "get user normal",
			path:       "/jwt/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.GET("/users", GetUsersController)
	req := httptest.NewRequest(http.MethodGet, "/jwt/users", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+token)

	e.ServeHTTP(rec, req)

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expectCode, rec.Code)
		body := rec.Body.String()

		var user UserResponses
		err := json.Unmarshal([]byte(body), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCase.sizeData, len(user.Users))
	}
}

func TestGetUserController(t *testing.T) {
	testCases := []Case{
		{
			name:       "success get user by id",
			path:       "/jwt/users/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed get user by id",
			path:       "/jwt/users/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}
	eJWT.GET("/users/:id", GetUserController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodGet, testCases[i].path, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)

	}
}

func TestDeleteUserController(t *testing.T) {
	testCases := []Case{
		{
			name:       "Success delete user",
			path:       "/jwt/users/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed delete user",
			path:       "/jwt/users/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}
	eJWT.DELETE("/users/:id", DeleteUserController)

	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodDelete, testCases[i].path, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)
	}
}

func TestUpdateUserController(t *testing.T) {
	testCases := []Case{
		{
			name:       "success update user",
			path:       "/jwt/users/1",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
		{
			name:       "failed update user",
			path:       "/jwt/users/2",
			expectCode: http.StatusBadRequest,
			sizeData:   1,
		},
	}

	payload := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     "fazri egi",
		Email:    "egi@gmail.com",
		Password: "egi123",
	}
	buf, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err.Error())
	}

	e := InitEchoTestAPI()
	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))

	InsertDataUserForTesting()
	token, err := Login("egi@gmail.com", "egi123", e)
	if err != nil {
		t.Fatal(err.Error())
	}

	eJWT.PUT("/users/:id", UpdateUserController)
	for i := 0; i < len(testCases); i++ {
		req := httptest.NewRequest(http.MethodPut, testCases[i].path, bytes.NewBuffer(buf))
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, testCases[i].expectCode, rec.Code)
	}
}

// func TestGetUserController(t *testing.T) {
// 	// create a new instance of echo
// 	e := echo.New()

// 	// create a new mock database
// 	mockDB := &MockDatabase{}

// 	// create a new test user
// 	testUser := &models.User{Name: "egi", Email: "egi@gmail.com", Password: "egi123"}

// 	// set up the mock database to return the test user when GetuserById is called
// 	mockDB.On("GetUserById", 1).Return(testUser, nil)

// 	// register the handler with the Echo instance
// 	e.GET("/users/:id", func(c echo.Context) error {
// 		// call the GetUserById function from the mock database
// 		id, _ := strconv.Atoi(c.Param("id"))
// 		user, err := mockDB.GetUserById(id)

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 				"message": err.Error(),
// 			})
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "success get user by id: " + c.Param("id"),
// 			"user":    user,
// 		})
// 	})

// 	// create a new HTTP GET request with a path that includes an "id" parameter
// 	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)

// 	// create a new HTTP recorder to capture the response
// 	rec := httptest.NewRecorder()

// 	// call the Echo handler with the request and recorder
// 	e.ServeHTTP(rec, req)

// 	// check that the response status code is 200 OK
// 	if rec.Code != http.StatusOK {
// 		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
// 	}

// 	// check that the response body matches the expected value
// 	// expectedBody := `{"message":"success get user by id: 123","user":{"id":"123","name":"John Doe"}}`
// 	// if rec.Body.String() != expectedBody {
// 	// 	t.Errorf("expected body %q but got %q", expectedBody, rec.Body.String())
// 	// }

// 	// assert that the GetUserById function was called with the correct parameter
// 	mockDB.AssertCalled(t, "GetUserById", 1)
// }

// // Define a mock database that implements the GetUserById method
// type MockDatabase struct {
// 	mock.Mock
// }

// func (m *MockDatabase) GetUserById(id interface{}) (*models.User, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.User), args.Error(1)
// }
