package helper

import "go.mongodb.org/mongo-driver/bson/primitive"

func ObjectIDFromHex(id string) primitive.ObjectID {
	objectId, err := primitive.ObjectIDFromHex(id)
	PanicIfError(err)
	return objectId
}
