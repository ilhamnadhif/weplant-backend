package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
)

type CustomerRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewCustomerRepository(collection *mongo.Collection) CustomerRepository {
	return &CustomerRepositoryImpl{
		Collection: collection,
	}
}

func (repository *CustomerRepositoryImpl) Create(ctx context.Context, customer schema.Customer) (schema.Customer, error) {
	res, err := repository.Collection.InsertOne(ctx, customer)
	if err != nil {
		return customer, err
	}
	customer.Id = res.InsertedID.(primitive.ObjectID)
	return customer, nil
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, customerId string) (schema.Customer, error) {
	objectId := helper.ObjectIDFromHex(customerId)
	var customer schema.Customer
	err := repository.Collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repository *CustomerRepositoryImpl) FindByEmail(ctx context.Context, email string) (schema.Customer, error) {
	var customer schema.Customer
	err := repository.Collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, customer schema.Customer) (schema.Customer, error) {
	_, err := repository.Collection.UpdateByID(ctx, customer.Id, bson.D{{"$set", customer}})
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, customerId string) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return err
	}
	return nil
}

// cart
func (repository *CustomerRepositoryImpl) PushProductToCart(ctx context.Context, customerId string, product schema.CartProduct) error {
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

func (repository *CustomerRepositoryImpl) UpdateProductQuantity(ctx context.Context, customerId string, product schema.CartProduct) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.UpdateOne(ctx, bson.D{
		{"$and", bson.A{
			bson.D{{"_id", objectId}},
			bson.D{{"carts.product_id", product.ProductId}},
			//bson.D{{"$where", fmt.Sprintf("for (let i = 0; i < this.carts.length; i++) {if (this.carts[i].product_id == '%s' && this.carts[i].quantity+%d >= 1) {return true} }", product.ProductId, product.Quantity)}},
		}},
	}, bson.D{
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

func (repository *CustomerRepositoryImpl) PullProductFromCart(ctx context.Context, customerId string, productId string) error {
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

func (repository *CustomerRepositoryImpl) PullProductFromAllCart(ctx context.Context, productId string) error {
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

// Transaction

func (repository *CustomerRepositoryImpl) CreateTransaction(ctx context.Context, customerId string, transaction schema.Transaction) error {
	objectId := helper.ObjectIDFromHex(customerId)
	_, err := repository.Collection.UpdateByID(ctx, objectId, bson.D{
		{"$push", bson.D{
			{"transactions", transaction},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}

func (repository *CustomerRepositoryImpl) DeleteTransaction(ctx context.Context, customerId string, transactionId string) error {
	objectCustomerId := helper.ObjectIDFromHex(customerId)
	objectTransactionId := helper.ObjectIDFromHex(transactionId)
	_, err := repository.Collection.UpdateByID(ctx, objectCustomerId, bson.D{
		{
			"$pull", bson.D{
				{
					"transactions", bson.D{
						{"_id", objectTransactionId},
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

func (repository *CustomerRepositoryImpl) CreateOrder(ctx context.Context, customerId string, order schema.OrderProduct) error {
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
