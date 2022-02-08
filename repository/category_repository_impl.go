package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
)

type CategoryRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewCategoryRepository(collection *mongo.Collection) CategoryRepository {
	return &CategoryRepositoryImpl{
		Collection: collection,
	}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	res, err := repository.Collection.InsertOne(ctx, category)
	if err != nil {
		return category, err
	}
	category.Id = res.InsertedID.(primitive.ObjectID)
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId string) (domain.Category, error) {
	var category domain.Category
	objectId := helper.ObjectIDFromHex(categoryId)
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&category)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	cursor, err := repository.Collection.Find(ctx, bson.D{})
	if err != nil {
		return categories, err
	}
	errorBind := cursor.All(ctx, &categories)
	if errorBind != nil {
		return categories, errorBind
	}
	return categories, nil
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	_, err := repository.Collection.UpdateByID(ctx, category.Id, bson.D{{"$set", category}})
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId string) error {
	objectId := helper.ObjectIDFromHex(categoryId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}
