package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductCategory struct {
	CategoryId string `bson:"category_id,omitempty"`
}

type Product struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   int                `bson:"created_at,omitempty"`
	UpdatedAt   int                `bson:"updated_at,omitempty"`
	MerchantId  string             `bson:"merchant_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       int                `bson:"price,omitempty"`
	Stock       int                `bson:"stock,omitempty"`
	MainImage   *Image             `bson:"main_image,omitempty"`
	Images      []*Image           `bson:"images,omitempty"`
	Categories  []*ProductCategory `bson:"categories,omitempty"`
}
