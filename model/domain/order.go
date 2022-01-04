package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderProduct struct {
	ProductId string `bson:"product_id,omitempty"`
	Price     int    `bson:"price,omitempty"`
	Quantity  int    `bson:"quantity,omitempty"`
}

type Order struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt int                `bson:"created_at,omitempty"`
	UpdatedAt int                `bson:"updated_at,omitempty"`
	Products  []*OrderProduct    `bson:"products,omitempty"`
	Address   *Address           `bson:"address,omitempty"`
}
