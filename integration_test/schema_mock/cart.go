package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/model/schema"
)

var CartProduct = schema.CartProduct{
	ProductId: primitive.NewObjectID().Hex(),
	Quantity:  15,
}
