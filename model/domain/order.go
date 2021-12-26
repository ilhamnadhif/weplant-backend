package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderProduct struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt int                `bson:"created_at,omitempty"`
	UpdatedAt int                `bson:"updated_at,omitempty"`
	ProductId string             `bson:"product_id,omitempty"`
	Price     int                `bson:"price,omitempty"`
	Quantity  int                `bson:"quantity,omitempty"`
	HasDone   *bool              `bson:"has_done,omitempty"`
	Address   *Address           `bson:"address,omitempty"`
}

type Order struct {
	Products []*OrderProduct `bson:"products,omitempty"`
}
