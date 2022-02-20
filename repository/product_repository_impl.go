package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

type ProductRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) ProductRepository {
	return &ProductRepositoryImpl{
		Collection: collection,
	}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, product schema.Product) (schema.Product, error) {
	res, err := repository.Collection.InsertOne(ctx, product)
	if err != nil {
		return product, err
	}
	product.Id = res.InsertedID.(primitive.ObjectID)
	return product, nil
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId string) (schema.Product, error) {
	var product schema.Product
	objectId := helper.ObjectIDFromHex(productId)
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) ([]schema.Product, error) {
	var products []schema.Product
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

func (repository *ProductRepositoryImpl) FindAllWithSearch(ctx context.Context, search string) ([]schema.Product, error) {
	var products []schema.Product
	cursor, err := repository.Collection.Find(ctx, bson.D{{"$and", bson.A{
		bson.D{{"$text", bson.D{{
			"$search", search,
		}}}},
	}}})
	if err != nil {
		return products, err
	}
	errorBind := cursor.All(ctx, &products)
	if errorBind != nil {
		return products, errorBind
	}
	return products, nil
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product schema.Product) (schema.Product, error) {
	_, err := repository.Collection.UpdateByID(ctx, product.Id, bson.D{{"$set", product}})
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) PushImageIntoImages(ctx context.Context, productId string, images []schema.Image) ([]schema.Image, error) {
	objectId := helper.ObjectIDFromHex(productId)
	_, err := repository.Collection.UpdateByID(ctx, objectId, bson.D{
		{
			"$push", bson.D{
				{
					"images", bson.D{
						{
							"$each", images,
						},
					},
				},
			},
		},
	})
	if err != nil {
		return images, err
	}
	return images, nil
}

func (repository *ProductRepositoryImpl) PullImageFromImages(ctx context.Context, productId string, imageId string) (schema.Image, error) {
	objectProductId := helper.ObjectIDFromHex(productId)
	objectImageId := helper.ObjectIDFromHex(imageId)

	var product schema.Product
	var image schema.Image

	err := repository.Collection.FindOneAndUpdate(ctx, bson.D{{"_id", objectProductId}}, bson.D{
		{
			"$pull", bson.D{
				{
					"images", bson.D{
						{"_id", objectImageId},
					},
				},
			},
		},
	}, options.FindOneAndUpdate().SetProjection(bson.D{
		{
			"images", bson.D{
				{
					"$elemMatch", bson.D{
						{"_id", objectImageId},
					},
				},
			},
		},
	})).Decode(&product)
	if err != nil {
		return image, err
	}

	if product.Images != nil {
		for _, img := range product.Images {
			if img.Id == objectImageId {
				image = img
			} else {
				return image, errors.New("Image Not Found")
			}
		}
	} else {
		return image, errors.New("Image Not Found")
	}
	return image, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, productId string) error {
	objectId := helper.ObjectIDFromHex(productId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

// merchant
func (repository *ProductRepositoryImpl) FindByMerchantId(ctx context.Context, merchantId string) ([]schema.Product, error) {
	var products []schema.Product
	cursor, err := repository.Collection.Find(ctx, bson.D{
		{"merchant_id", merchantId},
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

// category
func (repository *ProductRepositoryImpl) FindByCategoryId(ctx context.Context, categoryId string) ([]schema.Product, error) {
	var products []schema.Product
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

func (repository *ProductRepositoryImpl) PullCategoryIdFromProduct(ctx context.Context, categoryId string) error {
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

// transaction
func (repository *ProductRepositoryImpl) UpdateQuantity(ctx context.Context, product schema.Product) error {
	_, err := repository.Collection.UpdateByID(ctx, product.Id, bson.D{
		{
			"$inc", bson.D{
				{
					"stock", product.Stock,
				},
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
