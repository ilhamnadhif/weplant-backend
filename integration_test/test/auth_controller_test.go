package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"weplant-backend/integration_test/config"
	"weplant-backend/integration_test/schema_mock"
	"weplant-backend/model/web"
)

// Test Login Merchant

func TestLoginMerchant_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("FindByEmail", mock.Anything, mock.Anything).Return(schema_mock.Merchant, nil)

	router := config.SetupRouterTest()

	requestBody := web.LoginRequest{
		Email:    "ilham@gmail.com",
		Password: "12345",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/auth/merchant", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

}

func TestLoginMerchant_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("FindByEmail", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

	router := config.SetupRouterTest()

	requestBody := web.LoginRequest{
		Email:    "ilham@gmail.com",
		Password: "12345",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/auth/merchant", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
}

// Test Login Customer

func TestLoginCustomer_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindByEmail", mock.Anything, mock.Anything).Return(schema_mock.Customer, nil)

	router := config.SetupRouterTest()

	requestBody := web.LoginRequest{
		Email:    "ilham@gmail.com",
		Password: "12345",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/auth/customer", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLoginCustomer_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindByEmail", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

	router := config.SetupRouterTest()

	requestBody := web.LoginRequest{
		Email:    "ilham@gmail.com",
		Password: "12345",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/auth/customer", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
}
