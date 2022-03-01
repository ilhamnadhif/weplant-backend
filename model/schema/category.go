package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt int                `bson:"created_at,omitempty"`
	UpdatedAt int                `bson:"updated_at,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Slug      string             `bson:"slug"`
}
