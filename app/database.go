package app

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"weplant-backend/helper"
)

func GetConnection() *mongo.Client {
	mongoUri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	helper.PanicIfError(err)
	return client
}

func CloseConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	helper.PanicIfError(err)
}
