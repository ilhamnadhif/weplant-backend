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

var Transaction = schema.Transaction{
	Id:          primitive.NewObjectID(),
	CreatedAt:   helper.GetTimeNow(),
	UpdatedAt:   helper.GetTimeNow(),
	PaymentType: "gopay",
	Status:      "pending",
	QRCode:      "https://gqrodegopaty.com",
	Products: []schema.TransactionProduct{
		TransactionProduct,
		TransactionProduct,
	},
	Address: &Address,
}
