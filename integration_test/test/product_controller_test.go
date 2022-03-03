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
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
)

// Test FindById Product

func TestFindByIdProduct_Success(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/products/10", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestFindByIdProduct_Failed(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/products/10", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

// Test FindAll Product

func TestFindAllProduct_Success(t *testing.T) {
	config.ProductRepository.Mock.On("FindAll", context.Background(), mock.Anything, mock.Anything).Return([]schema.Product{
		schema_mock.Product,
		schema_mock.Product,
		schema_mock.Product,
	}, nil)
	config.ProductRepository.Mock.On("CountDocuments", context.Background()).Return(3, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/products", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestFindAllProduct_Failed(t *testing.T) {
	config.ProductRepository.Mock.On("FindAll", context.Background(), mock.Anything, mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("CountDocuments", context.Background()).Return(0, nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodGet, "https://test.com/api/v1/products", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
}

// Test Create Product

func TestCreateProduct_Success(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)
	config.ProductRepository.Mock.On("Create", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("merchant_id", primitive.NewObjectID().Hex())
	writer.WriteField("name", "toko ilham")
	writer.WriteField("description", "lorem dolor sit amet.")
	writer.WriteField("price", "50000")
	writer.WriteField("stock", "20")

	// main image
	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	// images
	file2, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file2.Write(uploadImageTest)
	writer.Close()

	file3, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file3.Write(uploadImageTest)
	writer.Close()

	// categories
	writer.WriteField("categories", primitive.NewObjectID().Hex())
	writer.WriteField("categories", primitive.NewObjectID().Hex())

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/products", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateProduct_Failed(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)
	config.ProductRepository.Mock.On("Create", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("merchant_id", primitive.NewObjectID().Hex())
	writer.WriteField("name", "toko ilham")
	writer.WriteField("description", "lorem dolor sit amet.")
	writer.WriteField("price", "50000")
	writer.WriteField("stock", "20")

	// main image
	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	// images
	file2, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file2.Write(uploadImageTest)
	writer.Close()

	file3, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file3.Write(uploadImageTest)
	writer.Close()

	// categories
	writer.WriteField("categories", primitive.NewObjectID().Hex())
	writer.WriteField("categories", primitive.NewObjectID().Hex())

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/products", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestCreateProduct_FailedUnauthorized(t *testing.T) {
	config.MerchantRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Merchant, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)
	config.ProductRepository.Mock.On("Create", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("merchant_id", primitive.NewObjectID().Hex())
	writer.WriteField("name", "toko ilham")
	writer.WriteField("description", "lorem dolor sit amet.")
	writer.WriteField("price", "50000")
	writer.WriteField("stock", "20")

	// main image
	file, err := writer.CreateFormFile("image", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file.Write(uploadImageTest)
	writer.Close()

	// images
	file2, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file2.Write(uploadImageTest)
	writer.Close()

	file3, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file3.Write(uploadImageTest)
	writer.Close()

	// categories
	writer.WriteField("categories", primitive.NewObjectID().Hex())
	writer.WriteField("categories", primitive.NewObjectID().Hex())

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/products", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test Update Product

func TestUpdateProduct_Success(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)
	config.ProductRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	requestBody := web.ProductUpdateRequest{
		Id:          primitive.NewObjectID().Hex(),
		UpdatedAt:   helper.GetTimeNow(),
		Name:        "bunga anggrek",
		Description: "bunga anggrek terbaik se indonesia",
		Price:       90000,
		Stock:       24,
		Categories: []web.ProductCategoryUpdateRequest{
			{
				CategoryId: primitive.NewObjectID().Hex(),
			},
			{
				CategoryId: primitive.NewObjectID().Hex(),
			},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "https://test.com/api/v1/products/10", bytes.NewReader(body))
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateProduct_Failed(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)
	config.ProductRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	requestBody := web.ProductUpdateRequest{
		Id:          primitive.NewObjectID().Hex(),
		UpdatedAt:   helper.GetTimeNow(),
		Name:        "bunga anggrek",
		Description: "bunga anggrek terbaik se indonesia",
		Price:       90000,
		Stock:       24,
		Categories: []web.ProductCategoryUpdateRequest{
			{
				CategoryId: primitive.NewObjectID().Hex(),
			},
			{
				CategoryId: primitive.NewObjectID().Hex(),
			},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "https://test.com/api/v1/products/10", bytes.NewReader(body))
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateProduct_FailedUnauthorized(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CategoryRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Category, nil)
	config.ProductRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Product, nil)

	router := config.SetupRouterTest()

	requestBody := web.ProductUpdateRequest{
		Id:          primitive.NewObjectID().Hex(),
		UpdatedAt:   helper.GetTimeNow(),
		Name:        "bunga anggrek",
		Description: "bunga anggrek terbaik se indonesia",
		Price:       90000,
		Stock:       24,
		Categories: []web.ProductCategoryUpdateRequest{
			{
				CategoryId: primitive.NewObjectID().Hex(),
			},
			{
				CategoryId: primitive.NewObjectID().Hex(),
			},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodPut, "https://test.com/api/v1/products/10", bytes.NewReader(body))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test UpdateMainImage Product

func TestUpdateMainImageProduct_Success(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.ProductRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
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

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/products/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateMainImageProduct_Failed(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.ProductRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
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

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/products/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateMainImageProduct_FailedUnauthorized(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.ProductRepository.Mock.On("Update", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
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

	request := httptest.NewRequest(http.MethodPatch, "https://test.com/api/v1/products/3/image", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test PushImageIntoImages Product

func TestPushImageIntoImagesProduct_Success(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.ProductRepository.Mock.On("PushImageIntoImages", context.Background(), mock.Anything, mock.Anything).Return([]schema.Image{
		schema_mock.Image,
		schema_mock.Image,
	}, nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// images
	file2, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file2.Write(uploadImageTest)
	writer.Close()

	file3, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file3.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/products/2/images", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestPushImageIntoImagesProduct_Failed(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.ProductRepository.Mock.On("PushImageIntoImages", context.Background(), mock.Anything, mock.Anything).Return([]schema.Image{
		schema_mock.Image,
		schema_mock.Image,
	}, nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// images
	file2, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file2.Write(uploadImageTest)
	writer.Close()

	file3, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file3.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/products/2/images", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestPushImageIntoImagesProduct_FailedUnauthorized(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CloudinaryRepository.Mock.On("UploadImage", context.Background(), mock.Anything, mock.Anything).Return("http://image.com/elonmusk.jpg", nil)
	config.ProductRepository.Mock.On("PushImageIntoImages", context.Background(), mock.Anything, mock.Anything).Return([]schema.Image{
		schema_mock.Image,
		schema_mock.Image,
	}, nil)

	router := config.SetupRouterTest()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// images
	file2, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file2.Write(uploadImageTest)
	writer.Close()

	file3, err := writer.CreateFormFile("images", "elonmusk.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	file3.Write(uploadImageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "https://test.com/api/v1/products/2/images", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test PullImageFromImages Product

func TestPullImageFromImagesProduct_Success(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.ProductRepository.Mock.On("PullImageFromImages", context.Background(), mock.Anything, mock.Anything).Return(schema_mock.Image, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/products/3/images/23", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestPullImageFromImagesProduct_Failed(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.ProductRepository.Mock.On("PullImageFromImages", context.Background(), mock.Anything, mock.Anything).Return(schema_mock.Image, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/products/3/images/23", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

}

func TestPullImageFromImagesProduct_FailedUnauthorized(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.ProductRepository.Mock.On("PullImageFromImages", context.Background(), mock.Anything, mock.Anything).Return(schema_mock.Image, nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/products/3/images/23", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}

// Test Delete Product

func TestDeleteProduct_Success(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PullProductFromAllCart", context.Background(), mock.Anything).Return(nil)
	config.ProductRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/products/4", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteProduct_Failed(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(nil, errors.New("error"))
	config.CustomerRepository.Mock.On("PullProductFromAllCart", context.Background(), mock.Anything).Return(nil)
	config.ProductRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/products/4", nil)
	request.Header.Add("Authorization", "Bearer "+config.GetJWTTokenTest("merchant"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
}

func TestDeleteProduct_FailedUnauthorized(t *testing.T) {
	config.ProductRepository.Mock.On("FindById", context.Background(), mock.Anything).Return(schema_mock.Product, nil)
	config.CustomerRepository.Mock.On("PullProductFromAllCart", context.Background(), mock.Anything).Return(nil)
	config.ProductRepository.Mock.On("Delete", context.Background(), mock.Anything).Return(nil)
	config.CloudinaryRepository.Mock.On("DeleteImage", context.Background(), mock.Anything).Return(nil)

	router := config.SetupRouterTest()

	request := httptest.NewRequest(http.MethodDelete, "https://test.com/api/v1/products/4", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
}
