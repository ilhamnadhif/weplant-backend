package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   int                `bson:"created_at,omitempty"`
	UpdatedAt   int                `bson:"updated_at,omitempty"`
	Address     *Address           `bson:"address,omitempty"`
}
