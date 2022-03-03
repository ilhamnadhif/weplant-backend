package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

var Customer = schema.Customer{
	Id:        primitive.NewObjectID(),
	CreatedAt: helper.GetTimeNow(),
	UpdatedAt: helper.GetTimeNow(),
	Email:     "ilham@gmail.com",
	Password:  "$2a$14$xV91BTRyTimTTfspZjepF.Wij3tcLO78HokFTyFr00ajQvoYmvKhe",
	UserName:  "ilham8725",
	Phone:     "081234567890",
	MainImage: &Image,
	Carts: []schema.CartProduct{
		CartProduct,
		CartProduct,
	},
	Transactions: []schema.Transaction{
		Transaction,
		Transaction,
	},
	Orders: []schema.OrderProduct{
		OrderProduct,
		OrderProduct,
	},
}
