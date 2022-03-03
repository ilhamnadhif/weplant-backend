package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

var TransactionProduct = schema.TransactionProduct{
	ProductId: primitive.NewObjectID().Hex(),
	Price:     20000,
	Quantity:  4,
}

var TransactionAction = schema.TransactionAction{
	Name:   "deeplink",
	Method: "GET",
	URL:    "https://gopay.com",
}

var Transaction = schema.Transaction{
	Id:          primitive.NewObjectID(),
	CreatedAt:   helper.GetTimeNow(),
	UpdatedAt:   helper.GetTimeNow(),
	PaymentType: "gopay",
	Status:      "pending",
	Actions: []schema.TransactionAction{
		TransactionAction,
		TransactionAction,
	},
	Products: []schema.TransactionProduct{
		TransactionProduct,
		TransactionProduct,
	},
	Address: &Address,
}
