package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
)

type MerchantRepository interface {
	Create(ctx context.Context, merchant domain.Merchant) (domain.Merchant, error)
	FindById(ctx context.Context, merchantId string) (domain.Merchant, error)
	Update(ctx context.Context, merchant domain.Merchant) (domain.Merchant, error)
	Delete(ctx context.Context, merchantId string) error

	// Manage Order
	PushProductToManageOrders(ctx context.Context, merchantId string, product domain.ManageOrderProduct) error
}

type merchantRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewMerchantRepository(collection *mongo.Collection) MerchantRepository {
	return &merchantRepositoryImpl{
		Collection: collection,
	}
}

func (repository *merchantRepositoryImpl) Create(ctx context.Context, merchant domain.Merchant) (domain.Merchant, error) {
	res, err := repository.Collection.InsertOne(ctx, merchant)
	if err != nil {
		return merchant, err
	}
	merchant.Id = res.InsertedID.(primitive.ObjectID)
	return merchant, nil
}

func (repository *merchantRepositoryImpl) FindById(ctx context.Context, merchantId string) (domain.Merchant, error) {
	var merchant domain.Merchant
	objectId := helper.ObjectIDFromHex(merchantId)
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&merchant)
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (repository *merchantRepositoryImpl) Update(ctx context.Context, merchant domain.Merchant) (domain.Merchant, error) {
	_, err := repository.Collection.UpdateByID(ctx, merchant.Id, bson.D{{"$set", merchant}})
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (repository *merchantRepositoryImpl) Delete(ctx context.Context, merchantId string) error {
	objectId := helper.ObjectIDFromHex(merchantId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

func (repository *merchantRepositoryImpl) PushProductToManageOrders(ctx context.Context, merchantId string, product domain.ManageOrderProduct) error {
	objectId := helper.ObjectIDFromHex(merchantId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"_id", objectId},
	}, bson.D{
		{"$push", bson.D{
			{"orders", product},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}
