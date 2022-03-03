package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"weplant-backend/helper"
	"weplant-backend/integration_test/config"
	"weplant-backend/integration_test/schema_mock"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
)

// Test Create Transaction

func TestCreateTransaction_Success(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CustomerRepository.Mock.On("PullProductFromCart", context.Background(), mock.Anything, mock.Anything).Return(nil)
	config.MidtransRepository.Mock.On("CreateTransaction", mock.Anything).Return(&coreapi.ChargeResponse{
		TransactionID: primitive.NewObjectID().Hex(),
		OrderID:       primitive.NewObjectID().Hex(),
		GrossAmount:   "200000",
		PaymentType:   "gopay",
	}, nil)
	config.CustomerRepository.Mock.On("CreateTransaction", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.AddressCreateRequest{
		Address:    "sudimoro",
		City:       "kudus",
		Province:   "jawa tengah",
		PostalCode: "679234",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/transactions/23", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateTransaction_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CustomerRepository.Mock.On("PullProductFromCart", context.Background(), mock.Anything, mock.Anything).Return(nil)
	config.MidtransRepository.Mock.On("CreateTransaction", mock.Anything).Return(&coreapi.ChargeResponse{
		TransactionID: primitive.NewObjectID().Hex(),
		OrderID:       primitive.NewObjectID().Hex(),
		GrossAmount:   "200000",
		PaymentType:   "gopay",
	}, nil)
	config.CustomerRepository.Mock.On("CreateTransaction", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.AddressCreateRequest{
		Address:    "sudimoro",
		City:       "kudus",
		Province:   "jawa tengah",
		PostalCode: "679234",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/transactions/23", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestCreateTransaction_FailedUnauthorized(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CustomerRepository.Mock.On("PullProductFromCart", context.Background(), mock.Anything, mock.Anything).Return(nil)
	config.MidtransRepository.Mock.On("CreateTransaction", mock.Anything).Return(&coreapi.ChargeResponse{
		TransactionID: primitive.NewObjectID().Hex(),
		OrderID:       primitive.NewObjectID().Hex(),
		GrossAmount:   "200000",
		PaymentType:   "gopay",
	}, nil)
	config.CustomerRepository.Mock.On("CreateTransaction", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := web.AddressCreateRequest{
		Address:    "sudimoro",
		City:       "kudus",
		Province:   "jawa tengah",
		PostalCode: "679234",
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/transactions/23", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test Cancel Transaction

func TestCancelTransaction_Success(t *testing.T) {
	customer := schema_mock.Customer
	customer.Transactions = []schema.Transaction{
		{
			Id:          helper.ObjectIDFromHex("621d9b2b5256a3aa8353dc08"),
			CreatedAt:   helper.GetTimeNow(),
			UpdatedAt:   helper.GetTimeNow(),
			PaymentType: "gopay",
			Status:      "pending",
			Actions: []schema.TransactionAction{
				schema_mock.TransactionAction,
				schema_mock.TransactionAction,
			},
			Products: []schema.TransactionProduct{
				schema_mock.TransactionProduct,
				schema_mock.TransactionProduct,
			},
			Address: &schema_mock.Address,
		},
	}
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(customer, nil)
	config.MidtransRepository.Mock.On("CancelTransaction", mock.Anything).Return(&coreapi.CancelResponse{}, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/transactions/23/transactions/621d9b2b5256a3aa8353dc08", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestCancelTransaction_Failed(t *testing.T) {
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.MidtransRepository.Mock.On("CancelTransaction", mock.Anything).Return(&coreapi.CancelResponse{}, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/transactions/23/transactions/621d9b2b5256a3aa8353dc08", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("customer"))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestCancelTransaction_FailedUnauthorized(t *testing.T) {
	customer := schema_mock.Customer
	customer.Transactions = []schema.Transaction{
		{
			Id:          helper.ObjectIDFromHex("621d9b2b5256a3aa8353dc08"),
			CreatedAt:   helper.GetTimeNow(),
			UpdatedAt:   helper.GetTimeNow(),
			PaymentType: "gopay",
			Status:      "pending",
			Actions: []schema.TransactionAction{
				schema_mock.TransactionAction,
				schema_mock.TransactionAction,
			},
			Products: []schema.TransactionProduct{
				schema_mock.TransactionProduct,
				schema_mock.TransactionProduct,
			},
			Address: &schema_mock.Address,
		},
	}
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(customer, nil)
	config.MidtransRepository.Mock.On("CancelTransaction", mock.Anything).Return(&coreapi.CancelResponse{}, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/transactions/23/transactions/621d9b2b5256a3aa8353dc08", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test Callback Transaction

func TestCallbackTransaction_Success(t *testing.T) {
	config.MidtransRepository.Mock.On("CheckTransaction", mock.Anything).Return(&coreapi.TransactionStatusResponse{
		OrderID:           primitive.NewObjectID().Hex(),
		PaymentType:       "gopay",
		TransactionID:     primitive.NewObjectID().Hex(),
		TransactionStatus: "success",
	}, nil)
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("CreateOrder", context.Background(), mock.Anything, mock.Anything).Return(nil)
	config.MerchantRepository.Mock.On("PushProductToManageOrders", context.Background(), mock.Anything, mock.Anything).Return(nil)
	config.ProductRepository.Mock.On("UpdateQuantity", context.Background(), mock.Anything).Return(nil)
	config.CustomerRepository.Mock.On("DeleteTransaction", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := `{
  "transaction_time": "2020-01-09 18:27:19",
  "transaction_status": "capture",
  "transaction_id": "57d5293c-e65f-4a29-95e4-5959c3fa335b",
  "status_message": "midtrans payment notification",
  "status_code": "200",
  "signature_key": "16d6f84b2fb0468e2a9cf99a8ac4e5d803d42180347aaa70cb2a7abb13b5c6130458ca9c71956a962c0827637cd3bc7d40b21a8ae9fab12c7c3efe351b18d00a",
  "payment_type": "credit_card",
  "order_id": "Postman-1578568851",
  "merchant_id": "G141532850",
  "masked_card": "481111-1114",
  "gross_amount": "10000.00",
  "fraud_status": "accept",
  "eci": "05",
  "currency": "IDR",
  "channel_response_message": "Approved",
  "channel_response_code": "00",
  "card_type": "credit",
  "bank": "bni",
  "approval_code": "1578569243927"
}`

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/callback", strings.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestCallbackTransaction_Failed(t *testing.T) {
	config.MidtransRepository.Mock.On("CheckTransaction", mock.Anything).Return(nil, &midtrans.Error{
		Message:        "error yaaaa",
		StatusCode:     500,
	})
	config.CustomerRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Customer, nil)
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("CreateOrder", context.Background(), mock.Anything, mock.Anything).Return(nil)
	config.MerchantRepository.Mock.On("PushProductToManageOrders", context.Background(), mock.Anything, mock.Anything).Return(nil)
	config.ProductRepository.Mock.On("UpdateQuantity", context.Background(), mock.Anything).Return(nil)
	config.CustomerRepository.Mock.On("DeleteTransaction", context.Background(), mock.Anything, mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	requestBody := `{
  "transaction_time": "2020-01-09 18:27:19",
  "transaction_status": "capture",
  "transaction_id": "57d5293c-e65f-4a29-95e4-5959c3fa335b",
  "status_message": "midtrans payment notification",
  "status_code": "200",
  "signature_key": "16d6f84b2fb0468e2a9cf99a8ac4e5d803d42180347aaa70cb2a7abb13b5c6130458ca9c71956a962c0827637cd3bc7d40b21a8ae9fab12c7c3efe351b18d00a",
  "payment_type": "credit_card",
  "order_id": "Postman-1578568851",
  "merchant_id": "G141532850",
  "masked_card": "481111-1114",
  "gross_amount": "10000.00",
  "fraud_status": "accept",
  "eci": "05",
  "currency": "IDR",
  "channel_response_message": "Approved",
  "channel_response_code": "00",
  "card_type": "credit",
  "bank": "bni",
  "approval_code": "1578569243927"
}`

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/callback", strings.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
}
