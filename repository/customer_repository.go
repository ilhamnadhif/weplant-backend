package repository

import (
	"context"
	"weplant-backend/model/schema"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer schema.Customer) (schema.Customer, error)
	FindById(ctx context.Context, customerId string) (schema.Customer, error)
	FindByEmail(ctx context.Context, email string) (schema.Customer, error)
	Update(ctx context.Context, customer schema.Customer) (schema.Customer, error)
	Delete(ctx context.Context, customerId string) error

	// Cart
	PushProductToCart(ctx context.Context, customerId string, product schema.CartProduct) error
	UpdateProductQuantity(ctx context.Context, customerId string, product schema.CartProduct) error
	PullProductFromCart(ctx context.Context, customerId string, productId string) error
	PullProductFromAllCart(ctx context.Context, productId string) error

	// Transaction
	CreateTransaction(ctx context.Context, customerId string, transaction schema.Transaction) error
	DeleteTransaction(ctx context.Context, customerId string, transactionId string) error

	// Order
	CreateOrder(ctx context.Context, customerId string, order schema.OrderProduct) error
}
