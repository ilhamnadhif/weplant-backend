package test_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"weplant-backend/integration_test/config"
	"weplant-backend/integration_test/schema_mock"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
)

// Test FindById Category

func TestFindByIdCategory_Success(t *testing.T) {
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)
	config.ProductRepository.Mock.On("FindByCategoryId", context.Background(), mock.Anything).Return([]schema.Product{
		schema_mock.Product,
		schema_mock.Product,
	}, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/categories/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestFindByIdCategory_Failed(t *testing.T) {
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/categories/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

// Test FindAll Category

func TestFindAllCategory_Success(t *testing.T) {
	config.CategoryRepository.Mock.On("FindAll", context.Background()).Return([]schema.Category{
		schema_mock.Category,
		schema_mock.Category,
	}, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/categories", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestFindAllCategory_Failed(t *testing.T) {
	config.CategoryRepository.Mock.On("FindAll", context.Background()).Return(nil, errors.New("error"))

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/categories", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

// Test Create Category

func TestCreateCategory_Success(t *testing.T) {
	config.CategoryRepository.Mock.On("Create", context.Background(), mock.Anything).Return(schema_mock.Category, nil)

	router := config.SetupRouterTest()

	requestBody := web.CategoryCreateRequest{
		Name: "sayuran",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/categories", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestCreateCategory_Failed(t *testing.T) {
	config.CategoryRepository.Mock.On("Create", context.Background(), mock.Anything).Return(schema_mock.Category, errors.New("error"))

	router := config.SetupRouterTest()

	requestBody := web.CategoryCreateRequest{
		Name: "sayuran",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/categories", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestCreateCategory_FailedUnauthorized(t *testing.T) {
	config.CategoryRepository.Mock.On("Create", context.Background(), mock.Anything).Return(schema_mock.Category, nil)

	router := config.SetupRouterTest()

	requestBody := web.CategoryCreateRequest{
		Name: "sayuran",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/categories", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}
