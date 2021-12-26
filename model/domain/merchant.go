package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Merchant struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt int                `bson:"created_at,omitempty"`
	UpdatedAt int                `bson:"updated_at,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Phone     string             `bson:"phone,omitempty"`
	Balance   int                `bson:"balance,omitempty"`
	MainImage *Image             `bson:"main_image,omitempty"`
	Address   *Address           `bson:"address,omitempty"`
}
