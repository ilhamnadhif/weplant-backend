package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	FindById(ctx context.Context, productId string) (domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, productId string) error
	PushImageIntoImages(ctx context.Context, productId string, image domain.Image) (domain.Image, error)
	FindByCategoryId(ctx context.Context, categoryId string) ([]domain.Product, error)
	PullCategoryIdFromProduct(ctx context.Context, categoryId string) error
}
type productRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) ProductRepository {
	return &productRepositoryImpl{
		Collection: collection,
	}
}

func (repository *productRepositoryImpl) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	res, err := repository.Collection.InsertOne(ctx, product)
	if err != nil {
		return product, err
	}
	product.Id = res.InsertedID.(primitive.ObjectID)
	return product, nil
}

func (repository *productRepositoryImpl) FindById(ctx context.Context, productId string) (domain.Product, error) {
	var product domain.Product
	objectId := helper.ObjectIDFromHex(productId)
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository *productRepositoryImpl) FindAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := repository.Collection.Find(ctx, bson.D{})
	if err != nil {
		return products, err
	}
	errorBind := cursor.All(ctx, &products)
	if errorBind != nil {
		return products, errorBind
	}
	return products, nil
}

func (repository *productRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	_, err := repository.Collection.UpdateByID(ctx, product.Id, bson.D{{"$set", product}})
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, productId string) error {
	objectId := helper.ObjectIDFromHex(productId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

func (repository *productRepositoryImpl) PushImageIntoImages(ctx context.Context, productId string, image domain.Image) (domain.Image, error) {
	objectId := helper.ObjectIDFromHex(productId)
	_, err := repository.Collection.UpdateByID(ctx, objectId, bson.D{
		{
			"$addToSet", bson.D{{
				"images", image,
			}},
		},
	})
	if err != nil {
		return image, err
	}
	return image, nil
}
func (repository *productRepositoryImpl) FindByCategoryId(ctx context.Context, categoryId string) ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := repository.Collection.Find(ctx, bson.D{
		{"categories", bson.D{
			{"$elemMatch", bson.D{
				{"category_id", categoryId},
			}},
		}},
	})
	if err != nil {
		return products, err
	}
	errBind := cursor.All(ctx, &products)
	if errBind != nil {
		return products, errBind
	}
	return products, nil
}

func (repository *productRepositoryImpl) PullCategoryIdFromProduct(ctx context.Context, categoryId string) error {
	_, err := repository.Collection.UpdateMany(ctx, bson.D{}, bson.D{
		{
			"$pull", bson.D{
				{
					"categories", bson.D{
						{"category_id", categoryId},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
