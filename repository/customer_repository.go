package repository

import (
	"context"
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

	// Transaction
	CreateTransaction(ctx context.Context, customerId string, transaction domain.Transaction) error
	DeleteTransaction(ctx context.Context, customerId string, transactionId string) error

	// Order
	CreateOrder(ctx context.Context, customerId string, order domain.OrderProduct) error
}
