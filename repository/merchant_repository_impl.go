package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

type MerchantRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewMerchantRepository(collection *mongo.Collection) MerchantRepository {
	return &MerchantRepositoryImpl{
		Collection: collection,
	}
}

func (repository *MerchantRepositoryImpl) Create(ctx context.Context, merchant schema.Merchant) (schema.Merchant, error) {
	res, err := repository.Collection.InsertOne(ctx, merchant)
	if err != nil {
		return merchant, err
	}
	merchant.Id = res.InsertedID.(primitive.ObjectID)
	return merchant, nil
}

func (repository *MerchantRepositoryImpl) FindById(ctx context.Context, merchantId string) (schema.Merchant, error) {
	var merchant schema.Merchant
	objectId := helper.ObjectIDFromHex(merchantId)
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&merchant)
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (repository *MerchantRepositoryImpl) FindByEmail(ctx context.Context, email string) (schema.Merchant, error) {
	var merchant schema.Merchant
	err := repository.Collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&merchant)
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (repository *MerchantRepositoryImpl) FindBySlug(ctx context.Context, slug string) (schema.Merchant, error) {
	var merchant schema.Merchant
	err := repository.Collection.FindOne(ctx, bson.D{{"slug", slug}}).Decode(&merchant)
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (repository *MerchantRepositoryImpl) Update(ctx context.Context, merchant schema.Merchant) (schema.Merchant, error) {
	_, err := repository.Collection.UpdateByID(ctx, merchant.Id, bson.D{{"$set", merchant}})
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

//func (repository *MerchantRepositoryImpl) UpdateBalance(ctx context.Context, merchant schema.Merchant) error {
//	_, err := repository.Collection.UpdateByID(ctx, merchant.Id, bson.D{
//		{
//			"$inc", bson.D{
//				{
//					"balance", merchant.Balance,
//				},
//			},
//		},
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (repository *MerchantRepositoryImpl) Delete(ctx context.Context, merchantId string) error {
	objectId := helper.ObjectIDFromHex(merchantId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

func (repository *MerchantRepositoryImpl) PushProductToManageOrders(ctx context.Context, merchantId string, product schema.ManageOrderProduct) error {
	objectId := helper.ObjectIDFromHex(merchantId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"_id", objectId},
	}, bson.D{
		{"$push", bson.D{
			{"orders", product},
		},
		},
		{
			"$inc", bson.D{
				{
					"balance", product.Price * product.Quantity,
				},
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
