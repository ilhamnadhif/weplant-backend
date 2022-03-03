package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

var ProductCategory = schema.ProductCategory{
	CategoryId: primitive.NewObjectID().Hex(),
}

var Product = schema.Product{
	Id:          primitive.NewObjectID(),
	CreatedAt:   helper.GetTimeNow(),
	UpdatedAt:   helper.GetTimeNow(),
	MerchantId:  primitive.NewObjectID().Hex(),
	Name:        "bunga melati",
	Slug:        "bunga-melati",
	Description: "lorem ipsum dolor sit amet",
	Price:       30000,
	Stock:       20,
	MainImage:   &Image,
	Images: []schema.Image{
		Image,
		Image,
	},
	Categories: []schema.ProductCategory{
		ProductCategory,
		ProductCategory,
	},
}
