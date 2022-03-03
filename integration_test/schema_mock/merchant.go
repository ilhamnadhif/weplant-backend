package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

var Merchant = schema.Merchant{
	Id:        primitive.NewObjectID(),
	CreatedAt: helper.GetTimeNow(),
	UpdatedAt: helper.GetTimeNow(),
	Email:     "ilham@gmail.com",
	Password:  "$2a$14$xV91BTRyTimTTfspZjepF.Wij3tcLO78HokFTyFr00ajQvoYmvKhe",
	Name:      "toko ilham",
	Slug:      "toko-ilham",
	Phone:     "081234567890",
	Balance:   2000000,
	MainImage: &Image,
	Orders: []schema.ManageOrderProduct{
		ManageOrderProduct,
		ManageOrderProduct,
	},
	Address: &Address,
}
