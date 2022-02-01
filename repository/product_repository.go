package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	FindById(ctx context.Context, productId string) (domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	PushImageIntoImages(ctx context.Context, productId string, images []domain.Image) ([]domain.Image, error)
	PullImageFromImages(ctx context.Context, productId string, imageId string) (domain.Image, error)
	Delete(ctx context.Context, productId string) error

	// merchant
	FindByMerchantId(ctx context.Context, merchantId string) ([]domain.Product, error)

	// category
	FindByCategoryId(ctx context.Context, categoryId string) ([]domain.Product, error)
	PullCategoryIdFromProduct(ctx context.Context, categoryId string) error

	// transaction
	UpdateQuantity(ctx context.Context, product domain.Product) error
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

func (repository *productRepositoryImpl) PushImageIntoImages(ctx context.Context, productId string, images []domain.Image) ([]domain.Image, error) {
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

func (repository *productRepositoryImpl) PullImageFromImages(ctx context.Context, productId string, imageId string) (domain.Image, error) {
	objectProductId := helper.ObjectIDFromHex(productId)
	objectImageId := helper.ObjectIDFromHex(imageId)

	var product domain.Product
	var image domain.Image

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
				image = *img
			} else {
				return image, errors.New("Image Not Found")
			}
		}
	} else {
		return image, errors.New("Image Not Found")
	}
	return image, nil
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, productId string) error {
	objectId := helper.ObjectIDFromHex(productId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

// merchant
func (repository *productRepositoryImpl) FindByMerchantId(ctx context.Context, merchantId string) ([]domain.Product, error) {
	var products []domain.Product
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

// transaction
func (repository *productRepositoryImpl) UpdateQuantity(ctx context.Context, product domain.Product) error {
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
