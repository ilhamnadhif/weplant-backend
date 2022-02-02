package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category domain.Category) (domain.Category, error)
	FindById(ctx context.Context, categoryId string) (domain.Category, error)
	FindAll(ctx context.Context) ([]domain.Category, error)
	Update(ctx context.Context, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, categoryId string) error
}

type categoryRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewCategoryRepository(collection *mongo.Collection) CategoryRepository {
	return &categoryRepositoryImpl{
		Collection: collection,
	}
}

func (repository *categoryRepositoryImpl) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	res, err := repository.Collection.InsertOne(ctx, category)
	if err != nil {
		return category, err
	}
	category.Id = res.InsertedID.(primitive.ObjectID)
	return category, nil
}

func (repository *categoryRepositoryImpl) FindById(ctx context.Context, categoryId string) (domain.Category, error) {
	var category domain.Category
	objectId := helper.ObjectIDFromHex(categoryId)
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&category)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository *categoryRepositoryImpl) FindAll(ctx context.Context) ([]domain.Category, error) {
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

func (repository *categoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	_, err := repository.Collection.UpdateByID(ctx, category.Id, bson.D{{"$set", category}})
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository *categoryRepositoryImpl) Delete(ctx context.Context, categoryId string) error {
	objectId := helper.ObjectIDFromHex(categoryId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}
