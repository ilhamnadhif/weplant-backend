package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

type AdminRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewAdminRepository(collection *mongo.Collection) AdminRepository {
	return &AdminRepositoryImpl{
		Collection: collection,
	}
}

func (repository *AdminRepositoryImpl) Create(ctx context.Context, admin schema.Admin) (schema.Admin, error) {
	res, err := repository.Collection.InsertOne(ctx, admin)
	if err != nil {
		return admin, err
	}
	admin.Id = res.InsertedID.(primitive.ObjectID)
	return admin, nil
}

func (repository *AdminRepositoryImpl) FindById(ctx context.Context, adminId string) (schema.Admin, error) {
	var admin schema.Admin
	objectId := helper.ObjectIDFromHex(adminId)
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&admin)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (repository *AdminRepositoryImpl) FindByEmail(ctx context.Context, email string) (schema.Admin, error) {
	var admin schema.Admin
	err := repository.Collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&admin)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (repository *AdminRepositoryImpl) FindAll(ctx context.Context) ([]schema.Admin, error) {
	var admins []schema.Admin
	cursor, err := repository.Collection.Find(ctx, bson.D{})
	if err != nil {
		return admins, err
	}
	err = cursor.All(ctx, &admins)
	if err != nil {
		return admins, err
	}
	return admins, nil
}

func (repository *AdminRepositoryImpl) Update(ctx context.Context, admin schema.Admin) (schema.Admin, error) {
	_, err := repository.Collection.UpdateByID(ctx, admin.Id, bson.D{{"$set", admin}})
	if err != nil {
		return admin, err
	}
	return admin, nil
}