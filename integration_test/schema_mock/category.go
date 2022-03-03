package schema_mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

var Category = schema.Category{
	Id:        primitive.NewObjectID(),
	CreatedAt: helper.GetTimeNow(),
	UpdatedAt: helper.GetTimeNow(),
	Name:      "sayuran",
	Slug:      "sayuran",
}
