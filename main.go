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

	// repository
	merchantRepository := repository.NewMerchantRepository(merchantCollection)
	cloudinaryRepository := repository.NewCloudinaryRepository(cloud)

	// service
	merchantService := service.NewMerchantService(merchantRepository, cloudinaryRepository)

	// controller
	merchantController := controller.NewMerchantController(merchantService)

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
	merchantRouter.PUT("/:merchantId", merchantController.UpdateById)
	merchantRouter.PATCH("/:merchantId/image", merchantController.UpdateMainImageById)
	merchantRouter.DELETE("/:merchantId", merchantController.DeleteById)

	errorRun := r.Run(":3000")
	if errorRun != nil {
		panic(errorRun)
	}
}
