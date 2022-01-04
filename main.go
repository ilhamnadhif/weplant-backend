package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"weplant-backend/config"
	"weplant-backend/controller"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/repository"
	"weplant-backend/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		helper.PanicIfError(err)
	}
	client := config.GetConnection()
	defer config.CloseConnection(client)
	database := client.Database("weplant-backend")

	// cloudinary get cloud
	cloud := config.GetCloud()

	// collection
	merchantCollection := database.Collection("merchant")
	merchantCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	categoryCollection := database.Collection("category")
	categoryCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	productCollection := database.Collection("product")
	productCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	customerCollection := database.Collection("customer")
	customerCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	// repository
	cloudinaryRepository := repository.NewCloudinaryRepository(cloud)
	merchantRepository := repository.NewMerchantRepository(merchantCollection)
	categoryRepository := repository.NewCategoryRepository(categoryCollection)
	productRepository := repository.NewProductRepository(productCollection)
	customerRepository := repository.NewCustomerRepository(customerCollection)
	midtransRepository := repository.NewMidtransRepository()

	// service
	merchantService := service.NewMerchantService(merchantRepository, cloudinaryRepository, productRepository)
	categoryService := service.NewCategoryService(categoryRepository, cloudinaryRepository, productRepository)
	productService := service.NewProductService(productRepository, cloudinaryRepository, categoryRepository, merchantRepository, customerRepository)
	customerService := service.NewCustomerService(customerRepository, productRepository, cloudinaryRepository)
	cartService := service.NewCartService(customerRepository, productRepository)
	orderService := service.NewOrderService(customerRepository, productRepository, midtransRepository, merchantRepository)

	// controller
	merchantController := controller.NewMerchantController(merchantService)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	customerController := controller.NewCustomerController(customerService)
	cartController := controller.NewCartController(cartService)
	orderController := controller.NewOrderController(orderService)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err,
		})
	}))

	// router
	v1 := r.Group("/api/v1")

	merchantRouter := v1.Group("/merchants")
	merchantRouter.POST("/", merchantController.Create)
	merchantRouter.GET("/:merchantId", merchantController.FindById)
	merchantRouter.GET("/:merchantId/orders", merchantController.FindManageOrderById)
	merchantRouter.PUT("/:merchantId", merchantController.Update)
	merchantRouter.PATCH("/:merchantId/image", merchantController.UpdateMainImage)
	merchantRouter.DELETE("/:merchantId", merchantController.Delete)

	categoryRouter := v1.Group("/categories")
	categoryRouter.POST("/", categoryController.Create)
	categoryRouter.GET("/:categoryId", categoryController.FindById)
	categoryRouter.GET("/", categoryController.FindAll)
	categoryRouter.PUT("/:categoryId", categoryController.Update)
	categoryRouter.PATCH("/:categoryId/image", categoryController.UpdateMainImage)
	categoryRouter.DELETE("/:categoryId", categoryController.Delete)

	productRouter := v1.Group("/products")
	productRouter.POST("/", productController.Create)
	productRouter.GET("/:productId", productController.FindById)
	productRouter.GET("/", productController.FindAll)
	productRouter.PUT("/:productId", productController.Update)
	productRouter.PATCH("/:productId/image", productController.UpdateMainImage)
	productRouter.POST("/:productId/images", productController.PushImageIntoImages)
	productRouter.DELETE("/:productId/images/:imageId", productController.PullImageFromImages)
	productRouter.DELETE("/:productId", productController.Delete)

	customerRouter := v1.Group("/customers")
	customerRouter.POST("/", customerController.Create)
	customerRouter.GET("/:customerId", customerController.FindById)
	customerRouter.GET("/:customerId/cart", customerController.FindCartById)
	customerRouter.GET("/:customerId/order", customerController.FindOrderById)
	customerRouter.PUT("/:customerId", customerController.Update)
	customerRouter.DELETE("/:customerId", customerController.Delete)

	cartRouter := v1.Group("/carts")
	cartRouter.POST("/:customerId", cartController.PushProductToCart)
	cartRouter.PATCH("/:customerId/products/:productId", cartController.UpdateProductQuantity)
	cartRouter.DELETE("/:customerId/products/:productId", cartController.PullProductFromCart)

	orderRouter := v1.Group("/orders")
	orderRouter.POST("/:customerId", orderController.CheckoutFromCart)
	orderRouter.POST("/callback", orderController.CallbackTransaction)

	errorRun := r.Run(":3000")
	if errorRun != nil {
		panic(errorRun)
	}
}
