package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt    int                `bson:"created_at,omitempty"`
	UpdatedAt    int                `bson:"updated_at,omitempty"`
	Email        string             `bson:"email,omitempty"`
	Password     string             `bson:"password,omitempty"`
	UserName     string             `bson:"user_name,omitempty"`
	Phone        string             `bson:"phone,omitempty"`
	Carts        []CartProduct      `bson:"carts,omitempty"`
	Transactions []Transaction      `bson:"transactions,omitempty"`
	Orders       []OrderProduct     `bson:"orders,omitempty"`
}
