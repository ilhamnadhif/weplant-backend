package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

var OrderProduct = schema.OrderProduct{
	Id:        primitive.NewObjectID(),
	CreatedAt: helper.GetTimeNow(),
	UpdatedAt: helper.GetTimeNow(),
	ProductId: primitive.NewObjectID().Hex(),
	Price:     30000,
	Quantity:  3,
	Address:   &Address,
}
