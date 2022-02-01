package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type TransactionProduct struct {
	ProductId string `bson:"product_id,omitempty"`
	Price     int    `bson:"price,omitempty"`
	Quantity  int    `bson:"quantity,omitempty"`
}

type Transaction struct {
	Id        primitive.ObjectID    `bson:"_id,omitempty"`
	CreatedAt int                   `bson:"created_at,omitempty"`
	UpdatedAt int                   `bson:"updated_at,omitempty"`
	Status    string                `json:"status"`
	QRCode    string                `json:"qr_code"`
	Products  []*TransactionProduct `bson:"products,omitempty"`
	Address   *Address              `bson:"address,omitempty"`
}
