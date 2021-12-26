package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	FindById(ctx context.Context, customerId string) (domain.Customer, error)
	Update(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Delete(ctx context.Context, customerId string) error

	PushProductToCart(ctx context.Context, customerId string, product domain.CartProduct) error
	UpdateProductQuantity(ctx context.Context, customerId string, product domain.CartProduct) error
	PullProductFromCart(ctx context.Context, customerId string, productId string) error
	PullProductFromAllCart(ctx context.Context, productId string) error

	CheckoutFromCart(ctx context.Context, customerId string, product domain.OrderProduct) error
}

type customerRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewCustomerRepository(collection *mongo.Collection) CustomerRepository {
	return &customerRepositoryImpl{
		Collection: collection,
	}
}

func (repository *customerRepositoryImpl) Create(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	res, err := repository.Collection.InsertOne(ctx, customer)
	if err != nil {
		return customer, err
	}
	customer.Id = res.InsertedID.(primitive.ObjectID)
	return customer, nil
}

func (repository *customerRepositoryImpl) FindById(ctx context.Context, customerId string) (domain.Customer, error) {
	objectId := helper.ObjectIDFromHex(customerId)
	var customer domain.Customer
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repository *customerRepositoryImpl) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	_, err := repository.Collection.UpdateByID(ctx, customer.Id, bson.D{{"$set", customer}})
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repository *customerRepositoryImpl) Delete(ctx context.Context, customerId string) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

//
func (repository *customerRepositoryImpl) PushProductToCart(ctx context.Context, customerId string, product domain.CartProduct) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"$and", bson.A{
			bson.D{{"_id", objectId}},
			bson.D{{"cart.products.product_id", bson.D{
				{"$ne", product.ProductId},
			}}},
		}},
	}, bson.D{
		{"$push", bson.D{
			{"cart.products", product},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}

func (repository *customerRepositoryImpl) UpdateProductQuantity(ctx context.Context, customerId string, product domain.CartProduct) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"$and", bson.A{
			bson.D{{"_id", objectId}},
			bson.D{{"cart.products.product_id", product.ProductId}},
			bson.D{{"$where", fmt.Sprintf("for (let i = 0; i < this.cart.products.length; i++) {if (this.cart.products[i].product_id == '%s' && this.cart.products[i].quantity+%d >= 1) {return true} }", product.ProductId, product.Quantity)}},
		}},
	}, bson.D{
		{
			"$set", bson.D{
				{"cart.products.$.updated_at", product.UpdatedAt},
			},
		},
		{"$inc", bson.D{
			{
				"cart.products.$.quantity", product.Quantity,
			},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}

func (repository *customerRepositoryImpl) PullProductFromCart(ctx context.Context, customerId string, productId string) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"_id", objectId},
	}, bson.D{
		{
			"$pull", bson.D{{
				"cart.products", bson.D{
					{"product_id", productId},
				},
			}},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (repository *customerRepositoryImpl) PullProductFromAllCart(ctx context.Context, productId string) error {
	_, err := repository.Collection.UpdateMany(ctx, bson.D{}, bson.D{
		{
			"$pull", bson.D{{
				"cart.products", bson.D{
					{"product_id", productId},
				},
			}},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (repository *customerRepositoryImpl) CheckoutFromCart(ctx context.Context, customerId string, product domain.OrderProduct) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"$and", bson.A{
			bson.D{{"_id", objectId}},
		}},
	}, bson.D{
		{"$push", bson.D{
			{"order.products", product},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}
