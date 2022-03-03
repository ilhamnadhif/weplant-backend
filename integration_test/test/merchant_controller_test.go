package test

import (
	"bytes"
	"context"
	_ "embed"
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
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
)

//go:embed "elonmusk.jpg"
var uploadImageTest []byte

// Test Create Merchant

func TestCreateMerchant_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("Create", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("email", "ilham@gmail.com")
	writer.WriteField("password", "12345")
	writer.WriteField("name", "toko ilham")
	writer.WriteField("phone", "081234567890")
	writer.WriteField("address", "Sudimoro")
	writer.WriteField("city", "Kudus")
	writer.WriteField("province", "Jawa Tengah")
	writer.WriteField("postal_code", "873623")

	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/merchants", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))

	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateMerchant_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("Create", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background()).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("email", "ilham@gmail.com")
	writer.WriteField("password", "12345")
	writer.WriteField("name", "toko ilham")
	writer.WriteField("phone", "081234567890")
	writer.WriteField("address", "Sudimoro")
	writer.WriteField("city", "Kudus")
	writer.WriteField("province", "Jawa Tengah")
	writer.WriteField("postal_code", "873623")

	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/merchants", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))

	assert.Equal(t, 500, response.StatusCode)
}

// Test FindById Merchant

func TestFindByIdMerchant_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)

	config.ProductRepository.Mock.On("FindByMerchantId", context.Background(), mock.Anything).Return([]schema.Product{
		schema_mock.Product,
		schema_mock.Product,
	}, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/merchants/10", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestFindByIdMerchant_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/merchants/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

// Test FindManageOrderById Merchant

func TestFindManageOrderByIdMerchant_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)

	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/merchants/10/orders", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestFindManageOrderByIdMerchant_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))

	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/merchants/10/orders", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestFindManageOrderByIdMerchant_FailedUnauthorized(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)

	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/merchants/10/orders", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

// Test Update Merchant

func TestUpdateMerchant_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.MerchantRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)

	router := config.SetupRouterTest()

	requestBody := web.MerchantUpdateRequest{
		Id:        primitive.NewObjectID().Hex(),
		UpdatedAt: helper.GetTimeNow(),
		Name:      "toko yanuar",
		Phone:     "098765432123",
		Address: &web.AddressUpdateRequest{
			Address:    "wonoketingal",
			City:       "kudus",
			Province:   "jawa tengah",
			PostalCode: "837454",
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "https://test.com/api/v1/merchants/10", bytes.NewReader(body))
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateMerchant_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.MerchantRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)

	router := config.SetupRouterTest()

	requestBody := web.MerchantUpdateRequest{
		Id:        primitive.NewObjectID().Hex(),
		UpdatedAt: helper.GetTimeNow(),
		Name:      "toko yanuar",
		Phone:     "098765432123",
		Address: &web.AddressUpdateRequest{
			Address:    "wonoketingal",
			City:       "kudus",
			Province:   "jawa tengah",
			PostalCode: "837454",
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "https://test.com/api/v1/merchants/10", bytes.NewReader(body))
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateMerchant_FailedUnauthorized(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.MerchantRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)

	router := config.SetupRouterTest()

	requestBody := web.MerchantUpdateRequest{
		Id:        primitive.NewObjectID().Hex(),
		UpdatedAt: helper.GetTimeNow(),
		Name:      "toko yanuar",
		Phone:     "098765432123",
		Address: &web.AddressUpdateRequest{
			Address:    "wonoketingal",
			City:       "kudus",
			Province:   "jawa tengah",
			PostalCode: "837454",
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "https://test.com/api/v1/merchants/10", bytes.NewReader(body))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test UpdateMainImage Merchant

func TestUpdateMainImage_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.MerchantRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
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

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/merchants/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateMainImage_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.MerchantRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
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

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/merchants/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateMainImage_FailedUnauthorized(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.MerchantRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
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

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/merchants/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test Delete Merchant

func TestDeleteMerchant_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.ProductRepository.Mock.On("FindByMerchantId", context.Background(), mock.Anything).Return(nil, nil)
	config.MerchantRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/merchants/10", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestDeleteMerchant_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("FindByMerchantId", context.Background(), mock.Anything).Return(nil, nil)
	config.MerchantRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/merchants/10", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

func TestDeleteMerchant_FailedUnauthorized(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.ProductRepository.Mock.On("FindByMerchantId", context.Background(), mock.Anything).Return(nil, nil)
	config.MerchantRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/merchants/10", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)

	//bytes, _ := io.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}
