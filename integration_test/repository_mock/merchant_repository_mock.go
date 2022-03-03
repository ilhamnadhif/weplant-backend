package repository_mock

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"weplant-backend/model/schema"
)

type MerchantRepositoryMock struct {
	Mock mock.Mock
}

func (repository *MerchantRepositoryMock) Create(ctx context.Context, merchant schema.Merchant) (schema.Merchant, error) {
	arguments := repository.Mock.Called(ctx, merchant)

	if arguments.Get(1) != nil {
		return schema.Merchant{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Merchant{}, errors.New("error")
	} else {
		merchant := arguments.Get(0).(schema.Merchant)
		return merchant, nil
	}
}

func (repository *MerchantRepositoryMock) FindById(ctx context.Context, merchantId string) (schema.Merchant, error) {
	arguments := repository.Mock.Called(ctx, merchantId)

	if arguments.Get(1) != nil {
		return schema.Merchant{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Merchant{}, errors.New("error")
	} else {
		merchant := arguments.Get(0).(schema.Merchant)
		return merchant, nil
	}
}

func (repository *MerchantRepositoryMock) FindByEmail(ctx context.Context, email string) (schema.Merchant, error) {
	arguments := repository.Mock.Called(ctx, email)

	if arguments.Get(1) != nil {
		return schema.Merchant{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Merchant{}, errors.New("error")
	} else {
		merchant := arguments.Get(0).(schema.Merchant)
		return merchant, nil
	}
}

func (repository *MerchantRepositoryMock) FindBySlug(ctx context.Context, slug string) (schema.Merchant, error) {

	arguments := repository.Mock.Called(ctx, slug)

	if arguments.Get(1) != nil {
		return schema.Merchant{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Merchant{}, errors.New("error")
	} else {
		merchant := arguments.Get(0).(schema.Merchant)
		return merchant, nil
	}

}

func (repository *MerchantRepositoryMock) Update(ctx context.Context, merchant schema.Merchant) (schema.Merchant, error) {

	arguments := repository.Mock.Called(ctx, merchant)

	if arguments.Get(1) != nil {
		return schema.Merchant{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Merchant{}, errors.New("error")
	} else {
		merchant := arguments.Get(0).(schema.Merchant)
		return merchant, nil
	}

}

func (repository *MerchantRepositoryMock) Delete(ctx context.Context, merchantId string) error {

	arguments := repository.Mock.Called(ctx, merchantId)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}

func (repository *MerchantRepositoryMock) PushProductToManageOrders(ctx context.Context, merchantId string, product schema.ManageOrderProduct) error {

	arguments := repository.Mock.Called(ctx, merchantId, product)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}
