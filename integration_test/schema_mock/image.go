package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/model/schema"
)

var Image = schema.Image{
	Id:       primitive.NewObjectID(),
	FileName: "elonmusk",
	URL:      "https://gambar.com/elonmusk.jpg",
}
