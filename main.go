package main

import (
	"context"
	"embed"
	_ "embed"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/fs"
	"net/http"
	"weplant-backend/config"
	"weplant-backend/controller"
	"weplant-backend/exception"
	"weplant-backend/helper"
	"weplant-backend/middleware"
	"weplant-backend/model/domain"
	"weplant-backend/repository"
	"weplant-backend/service"
)

//go:embed swagger
var spec embed.FS

func main() {

	swagger, err := fs.Sub(spec, "swagger")
	helper.PanicIfError(err)

	err = godotenv.Load()
	if err != nil {
		helper.PanicIfError(err)
	}
	client := config.GetConnection()
	defer config.CloseConnection(client)
	database := client.Database("weplant-backend")

	// cloudinary get cloud
	cloud := config.GetCloud()

	// get xendit key
	midtransKey := config.GetMidtransKey()

	// collection
	merchantCollection := database.Collection("merchant")
	merchantCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "slug", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	categoryCollection := database.Collection("category")
	categoryCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "slug", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	productCollection := database.Collection("product")
	productCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "slug", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	productCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{"name", "text"}},
	})
	customerCollection := database.Collection("customer")
	customerCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	adminCollection := database.Collection("admin")
	adminCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	// repository
	categoryRepository := repository.NewCategoryRepository(categoryCollection)
	merchantRepository := repository.NewMerchantRepository(merchantCollection)
	productRepository := repository.NewProductRepository(productCollection)
	customerRepository := repository.NewCustomerRepository(customerCollection)
	adminRepository := repository.NewAdminRepository(adminCollection)
	cloudinaryRepository := repository.NewCloudinaryRepository(cloud)
	midtransRepository := repository.NewMidtransRepository(midtransKey)

	// service
	categoryService := service.NewCategoryService(categoryRepository, cloudinaryRepository, productRepository)
	merchantService := service.NewMerchantService(merchantRepository, cloudinaryRepository, productRepository)
	productService := service.NewProductService(productRepository, cloudinaryRepository, categoryRepository, merchantRepository, customerRepository)
	customerService := service.NewCustomerService(customerRepository, productRepository, cloudinaryRepository)
	cartService := service.NewCartService(customerRepository, productRepository)
	transactionService := service.NewTransactionService(customerRepository, productRepository, midtransRepository, merchantRepository)
	jwtService := service.NewJWTService()
	authService := service.NewAuthService(merchantRepository, customerRepository, adminRepository)

	// controller
	categoryController := controller.NewCategoryController(categoryService)
	merchantController := controller.NewMerchantController(merchantService)
	productController := controller.NewProductController(productService)
	customerController := controller.NewCustomerController(customerService)
	cartController := controller.NewCartController(cartService)
	transactionController := controller.NewTransactionController(transactionService)
	authController := controller.NewAuthController(authService, jwtService)

	// middleware
	middleware2 := middleware.NewMiddleware(jwtService)

	// create admin
	res, err := adminRepository.FindAll(context.Background())
	helper.PanicIfError(err)
	if len(res) == 0 {
		_, err = adminRepository.Create(context.Background(), domain.Admin{
			Id:        primitive.ObjectID{},
			CreatedAt: helper.GetTimeNow(),
			UpdatedAt: helper.GetTimeNow(),
			Email:     "admin@admin.com",
			Password:  helper.HashPassword("admin999"),
		})
		helper.PanicIfError(err)
	}

	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler
	router.ServeFiles("/docs/*filepath", http.FS(swagger))

	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.POST("/api/v1/categories", middleware2.AuthMiddleware(categoryController.Create, "admin"))
	router.PUT("/api/v1/categories/:categoryId", middleware2.AuthMiddleware(categoryController.Update, "admin"))
	router.PATCH("/api/v1/categories/:categoryId/image", middleware2.AuthMiddleware(categoryController.UpdateMainImage, "admin"))
	router.DELETE("/api/v1/categories/:categoryId", middleware2.AuthMiddleware(categoryController.Delete, "admin"))

	router.POST("/api/v1/merchants", merchantController.Create)
	router.GET("/api/v1/merchants/:merchantId", merchantController.FindById)
	router.GET("/api/v1/merchants/:merchantId/orders", middleware2.AuthMiddleware(merchantController.FindManageOrderById, "merchant"))
	router.PUT("/api/v1/merchants/:merchantId", middleware2.AuthMiddleware(merchantController.Update, "merchant"))
	router.PATCH("/api/v1/merchants/:merchantId/image", middleware2.AuthMiddleware(merchantController.UpdateMainImage, "merchant"))
	router.DELETE("/api/v1/merchants/:merchantId", middleware2.AuthMiddleware(merchantController.Delete, "merchant"))

	router.GET("/api/v1/products/:productId", productController.FindById)
	router.GET("/api/v1/products", productController.FindAll)
	router.POST("/api/v1/products", middleware2.AuthMiddleware(productController.Create, "merchant"))
	router.PUT("/api/v1/products/:productId", middleware2.AuthMiddleware(productController.Update, "merchant"))
	router.PATCH("/api/v1/products/:productId/image", middleware2.AuthMiddleware(productController.UpdateMainImage, "merchant"))
	router.POST("/api/v1/products/:productId/images", middleware2.AuthMiddleware(productController.PushImageIntoImages, "merchant"))
	router.DELETE("/api/v1/products/:productId/images/:imageId", middleware2.AuthMiddleware(productController.PullImageFromImages, "merchant"))
	router.DELETE("/api/v1/products/:productId", middleware2.AuthMiddleware(productController.Delete, "merchant"))

	router.POST("/api/v1/customers/", customerController.Create)
	router.GET("/api/v1/customers/:customerId", customerController.FindById)
	router.GET("/api/v1/customers/:customerId/carts", middleware2.AuthMiddleware(customerController.FindCartById, "customer"))
	router.GET("/api/v1/customers/:customerId/transactions", middleware2.AuthMiddleware(customerController.FindTransactionById, "customer"))
	router.GET("/api/v1/customers/:customerId/orders", middleware2.AuthMiddleware(customerController.FindOrderById, "customer"))
	router.PUT("/api/v1/customers/:customerId", middleware2.AuthMiddleware(customerController.Update, "customer"))
	router.DELETE("/api/v1/customers/:customerId", middleware2.AuthMiddleware(customerController.Delete, "customer"))

	router.POST("/api/v1/carts/:customerId", middleware2.AuthMiddleware(cartController.PushProductToCart, "customer"))
	router.PATCH("/api/v1/carts/:customerId/products/:productId", middleware2.AuthMiddleware(cartController.UpdateProductQuantity, "customer"))
	router.DELETE("/api/v1/carts/:customerId/products/:productId", middleware2.AuthMiddleware(cartController.PullProductFromCart, "customer"))

	router.POST("/api/v1/callback", transactionController.Callback)
	router.POST("/api/v1/transactions/:customerId", middleware2.AuthMiddleware(transactionController.Create, "customer"))
	router.DELETE("/api/v1/transactions/:customerId/transactions/:transactionId", middleware2.AuthMiddleware(transactionController.Cancel, "customer"))

	router.POST("/api/v1/auth/merchant", authController.LoginMerchant)
	router.POST("/api/v1/auth/customer", authController.LoginCustomer)
	router.POST("/api/v1/auth/admin", authController.LoginAdmin)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
