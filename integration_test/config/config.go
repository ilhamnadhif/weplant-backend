package config

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/mock"
	"weplant-backend/app"
	"weplant-backend/controller"
	"weplant-backend/integration_test/repository_mock"
	"weplant-backend/model/web"
	"weplant-backend/pkg"
	"weplant-backend/service"
)

var MerchantRepository = repository_mock.MerchantRepositoryMock{Mock: mock.Mock{}}
var ProductRepository = repository_mock.ProductRepositoryMock{Mock: mock.Mock{}}
var CategoryRepository = repository_mock.CategoryRepositoryMock{Mock: mock.Mock{}}
var CustomerRepository = repository_mock.CustomerRepositoryMock{Mock: mock.Mock{}}
var CloudinaryRepository = repository_mock.CloudinaryRepositoryMock{Mock: mock.Mock{}}
var MidtransRepository = repository_mock.MidtransRepositoryMock{Mock: mock.Mock{}}

func SetupRouterTest() *httprouter.Router {
	// service
	authService := service.NewAuthService(&MerchantRepository, &CustomerRepository)
	merchantService := service.NewMerchantService(&MerchantRepository, &CloudinaryRepository, &ProductRepository)
	productService := service.NewProductService(&ProductRepository, &CloudinaryRepository, &CategoryRepository, &MerchantRepository, &CustomerRepository)
	categoryService := service.NewCategoryService(&CategoryRepository, &ProductRepository)
	customerService := service.NewCustomerService(&CustomerRepository, &ProductRepository, &CloudinaryRepository)
	cartService := service.NewCartService(&CustomerRepository, &ProductRepository)
	transactionService := service.NewTransactionService(&CustomerRepository, &ProductRepository, &MidtransRepository, &MerchantRepository)

	// controller
	authController := controller.NewAuthController(authService)
	merchantController := controller.NewMerchantController(merchantService)
	productController := controller.NewProductController(productService)
	categoryController := controller.NewCategoryController(categoryService)
	customerController := controller.NewCustomerController(customerService)
	cartController := controller.NewCartController(cartService)
	transactionController := controller.NewTransactionController(transactionService)

	router := app.NewRouter(nil, authController, merchantController, productController, categoryController, customerController, cartController, transactionController)

	return router
}

func GetJWTTokenTest(role string) string {
	return pkg.GenerateToken(web.JWTPayload{
		Id:   "1",
		Role: role,
	})
}
