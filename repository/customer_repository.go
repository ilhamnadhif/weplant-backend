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
	FindByEmail(ctx context.Context, email string) (domain.Customer, error)
	Update(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Delete(ctx context.Context, customerId string) error

	// Cart
	PushProductToCart(ctx context.Context, customerId string, product domain.CartProduct) error
	UpdateProductQuantity(ctx context.Context, customerId string, product domain.CartProduct) error
	PullProductFromCart(ctx context.Context, customerId string, productId string) error
	PullProductFromAllCart(ctx context.Context, productId string) error

	// Order
	CheckoutFromCart(ctx context.Context, customerId string, order domain.Order) error
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

func (repository *customerRepositoryImpl) FindByEmail(ctx context.Context, email string) (domain.Customer, error) {
	var customer domain.Customer
	err := repository.Collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&customer)
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
			bson.D{{"carts.product_id", bson.D{
				{"$ne", product.ProductId},
			}}},
		}},
	}, bson.D{
		{"$push", bson.D{
			{"carts", product},
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
			bson.D{{"carts.product_id", product.ProductId}},
			bson.D{{"$where", fmt.Sprintf("for (let i = 0; i < this.carts.length; i++) {if (this.carts[i].product_id == '%s' && this.carts[i].quantity+%d >= 1) {return true} }", product.ProductId, product.Quantity)}},
		}},
	}, bson.D{
		{
			"$set", bson.D{
				{"carts.$.updated_at", product.UpdatedAt},
			},
		},
		{"$inc", bson.D{
			{
				"carts.$.quantity", product.Quantity,
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
				"carts", bson.D{
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
				"carts", bson.D{
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

func (repository *customerRepositoryImpl) CheckoutFromCart(ctx context.Context, customerId string, order domain.Order) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"$and", bson.A{
			bson.D{{"_id", objectId}},
		}},
	}, bson.D{
		{"$push", bson.D{
			{"orders", order},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}
