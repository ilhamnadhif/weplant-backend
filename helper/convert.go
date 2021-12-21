package helper

import "go.mongodb.org/mongo-driver/bson/primitive"

func ObjectIDFromHex(id string) primitive.ObjectID {
	objectId, errorId := primitive.ObjectIDFromHex(id)
	PanicIfError(errorId)
	return objectId
}
