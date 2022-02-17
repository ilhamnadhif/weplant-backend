package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	FileName string             `bson:"file_name,omitempty"`
	URL      string             `bson:"url,omitempty"`
}
