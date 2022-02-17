package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type TransactionProduct struct {
	ProductId string `bson:"product_id,omitempty"`
	Price     int    `bson:"price,omitempty"`
	Quantity  int    `bson:"quantity,omitempty"`
}

type TransactionAction struct {
	Name   string `bson:"name,omitempty"`
	Method string `bson:"method,omitempty"`
	URL    string `bson:"url,omitempty"`
}

type Transaction struct {
	Id          primitive.ObjectID   `bson:"_id,omitempty"`
	CreatedAt   int                  `bson:"created_at,omitempty"`
	UpdatedAt   int                  `bson:"updated_at,omitempty"`
	PaymentType string               `bson:"payment_type,omitempty"`
	Status      string               `bson:"status,omitempty"`
	Actions     []TransactionAction  `bson:"actions,omitempty"`
	Products    []TransactionProduct `bson:"products,omitempty"`
	Address     *Address             `bson:"address,omitempty"`
}
