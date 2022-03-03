package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"weplant-backend/integration_test/config"
	"weplant-backend/integration_test/schema_mock"
	"weplant-backend/model/web"
)

// Test PushProductToCart Cart

func TestPushProductToCartCart_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PushProductToCart", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.CartProductCreateRequest{
		CustomerId: primitive.NewObjectID().Hex(),
		ProductId:  primitive.NewObjectID().Hex(),
		Quantity:   1,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/carts/121", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestPushProductToCartCart_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PushProductToCart", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.CartProductCreateRequest{
		CustomerId: primitive.NewObjectID().Hex(),
		ProductId:  primitive.NewObjectID().Hex(),
		Quantity:   1,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/carts/121", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestPushProductToCartCart_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PushProductToCart", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.CartProductCreateRequest{
		CustomerId: primitive.NewObjectID().Hex(),
		ProductId:  primitive.NewObjectID().Hex(),
		Quantity:   1,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/carts/121", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test UpdateProductQuantity Cart

func TestUpdateProductQuantityCart_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("UpdateProductQuantity", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.CartProductUpdateRequest{
		CustomerId: primitive.NewObjectID().Hex(),
		ProductId:  primitive.NewObjectID().Hex(),
		Quantity:   4,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8080/api/v1/carts/121/products/12", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateProductQuantityCart_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("UpdateProductQuantity", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.CartProductUpdateRequest{
		CustomerId: primitive.NewObjectID().Hex(),
		ProductId:  primitive.NewObjectID().Hex(),
		Quantity:   4,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8080/api/v1/carts/121/products/12", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateProductQuantityCart_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("UpdateProductQuantity", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.CartProductUpdateRequest{
		CustomerId: primitive.NewObjectID().Hex(),
		ProductId:  primitive.NewObjectID().Hex(),
		Quantity:   4,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8080/api/v1/carts/121/products/12", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test PullProductFromCart Cart

func TestPullProductFromCartCart_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PullProductFromCart", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/carts/121/products/12", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestPullProductFromCartCart_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PullProductFromCart", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/carts/121/products/12", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestPullProductFromCartCart_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PullProductFromCart", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/carts/121/products/12", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}
