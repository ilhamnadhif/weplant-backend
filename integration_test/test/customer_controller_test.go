package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"weplant-backend/helper"
	"weplant-backend/integration_test/config"
	"weplant-backend/integration_test/schema_mock"
	"weplant-backend/model/web"
)

// Test Create Customer

func TestCreateCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("Create", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)

	router := config.SetupRouterTest()

	requestBody := web.CustomerCreateRequest{
		CreatedAt: helper.GetTimeNow(),
		UpdatedAt: helper.GetTimeNow(),
		Email:     "ilham@gmail.com",
		Password:  "12345",
		UserName:  "ilham8725",
		Phone:     "081234567890",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/customers", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestCreateCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("Create", context.Background(), mock.Anything).Return(nil, errors.New("error"))

	router := config.SetupRouterTest()

	requestBody := web.CustomerCreateRequest{
		CreatedAt: helper.GetTimeNow(),
		UpdatedAt: helper.GetTimeNow(),
		Email:     "ilham@gmail.com",
		Password:  "12345",
		UserName:  "ilham8725",
		Phone:     "081234567890",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/customers", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
}

// Test FindById Customer

func TestFindByIdCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestFindByIdCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

// Test FindCartById Customer

func TestFindCartByIdCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/carts", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestFindCartByIdCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/carts", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

}

func TestFindCartByIdCustomer_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/carts", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test FindTransactionById Customer

func TestFindTransactionByIdCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/transactions", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestFindTransactionByIdCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/transactions", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestFindTransactionByIdCustomer_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/transactions", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test FindOrderById Customer

func TestFindOrderByIdCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/orders", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestFindOrderByIdCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/orders", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestFindOrderByIdCustomer_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/customers/34/orders", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test Update Customer

func TestUpdateCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CustomerRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)

	router := config.SetupRouterTest()

	requestBody := web.CustomerUpdateRequest{
		Id:        primitive.NewObjectID().Hex(),
		UpdatedAt: helper.GetTimeNow(),
		UserName:  "yanuarnauval",
		Phone:     "098765432123",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/customers/34", bytes.NewReader(body))
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CustomerRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)

	router := config.SetupRouterTest()

	requestBody := web.CustomerUpdateRequest{
		Id:        primitive.NewObjectID().Hex(),
		UpdatedAt: helper.GetTimeNow(),
		UserName:  "yanuarnauval",
		Phone:     "098765432123",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/customers/34", bytes.NewReader(body))
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateCustomer_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CustomerRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)

	router := config.SetupRouterTest()

	requestBody := web.CustomerUpdateRequest{
		Id:        primitive.NewObjectID().Hex(),
		UpdatedAt: helper.GetTimeNow(),
		UserName:  "yanuarnauval",
		Phone:     "098765432123",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/customers/34", bytes.NewReader(body))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test UpdateMainImage Customer

func TestUpdateMainImageCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CustomerRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/customers/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateMainImageCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CustomerRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/customers/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateMainImageCustomer_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CustomerRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/customers/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test Delete Customer

func TestDeleteCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CustomerRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/customers/10", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CustomerRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/customers/10", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestDeleteCustomer_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.CustomerRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/customers/10", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}
