package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"weplant-backend/config"
	"weplant-backend/controller"
	"weplant-backend/helper"
	"weplant-backend/middleware"
	"weplant-backend/model/domain"
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
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

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

	categoryRouter := v1.Group("/categories")
	categoryRouter.GET("/:categoryId", categoryController.FindById)
	categoryRouter.GET("/", categoryController.FindAll)
	categoryRouter.Use(authMiddleware.AuthJWT("admin"))
	categoryRouter.POST("/", categoryController.Create)
	categoryRouter.PUT("/:categoryId", categoryController.Update)
	categoryRouter.PATCH("/:categoryId/image", categoryController.UpdateMainImage)
	categoryRouter.DELETE("/:categoryId", categoryController.Delete)

	merchantRouter := v1.Group("/merchants")
	merchantRouter.POST("/", merchantController.Create)
	merchantRouter.GET("/:merchantId", merchantController.FindById)
	merchantRouter.Use(authMiddleware.AuthJWT("merchant"))
	merchantRouter.GET("/:merchantId/orders", merchantController.FindManageOrderById)
	merchantRouter.PUT("/:merchantId", merchantController.Update)
	merchantRouter.PATCH("/:merchantId/image", merchantController.UpdateMainImage)
	merchantRouter.DELETE("/:merchantId", merchantController.Delete)

	productRouter := v1.Group("/products")
	productRouter.GET("/:productId", productController.FindById)
	productRouter.GET("/", productController.FindAll)
	productRouter.Use(authMiddleware.AuthJWT("merchant"))
	productRouter.POST("/", productController.Create)
	productRouter.PUT("/:productId", productController.Update)
	productRouter.PATCH("/:productId/image", productController.UpdateMainImage)
	productRouter.POST("/:productId/images", productController.PushImageIntoImages)
	productRouter.DELETE("/:productId/images/:imageId", productController.PullImageFromImages)
	productRouter.DELETE("/:productId", productController.Delete)

	customerRouter := v1.Group("/customers")
	customerRouter.POST("/", customerController.Create)
	customerRouter.GET("/:customerId", customerController.FindById)
	customerRouter.Use(authMiddleware.AuthJWT("customer"))
	customerRouter.GET("/:customerId/carts", customerController.FindCartById)
	customerRouter.GET("/:customerId/transactions", customerController.FindTransactionById)
	customerRouter.GET("/:customerId/orders", customerController.FindOrderById)
	customerRouter.PUT("/:customerId", customerController.Update)
	customerRouter.DELETE("/:customerId", customerController.Delete)

	cartRouter := v1.Group("/carts")
	cartRouter.Use(authMiddleware.AuthJWT("customer"))
	cartRouter.POST("/:customerId", cartController.PushProductToCart)
	cartRouter.PATCH("/:customerId/products/:productId", cartController.UpdateProductQuantity)
	cartRouter.DELETE("/:customerId/products/:productId", cartController.PullProductFromCart)

	transactionRouter := v1.Group("/transactions")
	transactionRouter.POST("/callback", transactionController.Callback)
	transactionRouter.Use(authMiddleware.AuthJWT("customer"))
	transactionRouter.POST("/:customerId", transactionController.Create)
	transactionRouter.DELETE("/:customerId/transactions/:transactionId", transactionController.Cancel)

	authRouter := v1.Group("/auth")
	authRouter.POST("/merchant", authController.LoginMerchant)
	authRouter.POST("/customer", authController.LoginCustomer)
	authRouter.POST("/admin", authController.LoginAdmin)

	errorRun := r.Run(":3000")
	if errorRun != nil {
		panic(errorRun)
	}
}
