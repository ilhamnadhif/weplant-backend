package main

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/fs"
	"net/http"
	"os"
	"weplant-backend/app"
	"weplant-backend/controller"
	"weplant-backend/helper"
	"weplant-backend/pkg"
	"weplant-backend/repository"
	"weplant-backend/service"
)

//go:embed swagger
var spec embed.FS

func main() {

	swagger, err := fs.Sub(spec, "swagger")
	helper.PanicIfError(err)

	pkg.GoDotENV()

	client := app.GetConnection()
	defer app.CloseConnection(client)
	database := client.Database("weplant-backend")

	// cloudinary get cloud
	cloud := app.GetCloud()

	// get xendit key
	midtransKey := app.GetMidtransKey()

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
	productCollection := database.Collection("product")
	productCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "slug", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	productCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{"name", "text"}},
	})
	categoryCollection := database.Collection("category")
	categoryCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "slug", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	customerCollection := database.Collection("customer")
	customerCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	// repository
	merchantRepository := repository.NewMerchantRepository(merchantCollection)
	productRepository := repository.NewProductRepository(productCollection)
	categoryRepository := repository.NewCategoryRepository(categoryCollection)
	customerRepository := repository.NewCustomerRepository(customerCollection)
	cloudinaryRepository := repository.NewCloudinaryRepository(cloud)
	midtransRepository := repository.NewMidtransRepository(midtransKey)

	// service
	authService := service.NewAuthService(merchantRepository, customerRepository)
	merchantService := service.NewMerchantService(merchantRepository, cloudinaryRepository, productRepository)
	productService := service.NewProductService(productRepository, cloudinaryRepository, categoryRepository, merchantRepository, customerRepository)
	categoryService := service.NewCategoryService(categoryRepository, productRepository)
	customerService := service.NewCustomerService(customerRepository, productRepository, cloudinaryRepository)
	cartService := service.NewCartService(customerRepository, productRepository)
	transactionService := service.NewTransactionService(customerRepository, productRepository, midtransRepository, merchantRepository)

	// controller
	authController := controller.NewAuthController(authService)
	merchantController := controller.NewMerchantController(merchantService)
	productController := controller.NewProductController(productService)
	categoryController := controller.NewCategoryController(categoryService)
	customerController := controller.NewCustomerController(customerService)
	cartController := controller.NewCartController(cartService)
	transactionController := controller.NewTransactionController(transactionService)

	router := app.NewRouter(swagger, authController, merchantController, productController, categoryController, customerController, cartController, transactionController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(fmt.Sprintf("app listening on port %s", port))

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
