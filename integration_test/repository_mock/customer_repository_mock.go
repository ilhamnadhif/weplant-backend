package repository_mock

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"weplant-backend/model/schema"
)

type CustomerRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CustomerRepositoryMock) Create(ctx context.Context, customer schema.Customer) (schema.Customer, error) {

	arguments := repository.Mock.Called(ctx, customer)

	if arguments.Get(1) != nil {
		return schema.Customer{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Customer{}, errors.New("error")
	} else {
		customer := arguments.Get(0).(schema.Customer)
		return customer, nil
	}

}

func (repository *CustomerRepositoryMock) FindById(ctx context.Context, customerId string) (schema.Customer, error) {
	arguments := repository.Mock.Called(ctx, customerId)

	if arguments.Get(1) != nil {
		return schema.Customer{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Customer{}, errors.New("error")
	} else {
		customer := arguments.Get(0).(schema.Customer)
		return customer, nil
	}
}

func (repository *CustomerRepositoryMock) FindByEmail(ctx context.Context, email string) (schema.Customer, error) {
	arguments := repository.Mock.Called(ctx, email)
	if arguments.Get(1) != nil {
		return schema.Customer{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Customer{}, errors.New("error")
	} else {
		return arguments.Get(0).(schema.Customer), nil
	}
}

func (repository *CustomerRepositoryMock) Update(ctx context.Context, customer schema.Customer) (schema.Customer, error) {

	arguments := repository.Mock.Called(ctx, customer)

	if arguments.Get(1) != nil {
		return schema.Customer{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Customer{}, errors.New("error")
	} else {
		customer := arguments.Get(0).(schema.Customer)
		return customer, nil
	}

}

func (repository *CustomerRepositoryMock) Delete(ctx context.Context, customerId string) error {

	arguments := repository.Mock.Called(ctx, customerId)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}
}

func (repository *CustomerRepositoryMock) PushProductToCart(ctx context.Context, customerId string, product schema.CartProduct) error {

	arguments := repository.Mock.Called(ctx, customerId, product)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}

func (repository *CustomerRepositoryMock) UpdateProductQuantity(ctx context.Context, customerId string, product schema.CartProduct) error {

	arguments := repository.Mock.Called(ctx, customerId, product)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}
}

func (repository *CustomerRepositoryMock) PullProductFromCart(ctx context.Context, customerId string, productId string) error {

	arguments := repository.Mock.Called(ctx, customerId, productId)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}

func (repository *CustomerRepositoryMock) PullProductFromAllCart(ctx context.Context, productId string) error {

	arguments := repository.Mock.Called(ctx, productId)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}

func (repository *CustomerRepositoryMock) CreateTransaction(ctx context.Context, customerId string, transaction schema.Transaction) error {

	arguments := repository.Mock.Called(ctx, customerId, transaction)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}

func (repository *CustomerRepositoryMock) DeleteTransaction(ctx context.Context, customerId string, transactionId string) error {

	arguments := repository.Mock.Called(ctx, customerId, transactionId)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}

func (repository *CustomerRepositoryMock) CreateOrder(ctx context.Context, customerId string, order schema.OrderProduct) error {

	arguments := repository.Mock.Called(ctx, customerId, order)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}
