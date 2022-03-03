package repository_mock

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"weplant-backend/model/schema"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) Create(ctx context.Context, product schema.Product) (schema.Product, error) {
	arguments := repository.Mock.Called(ctx, product)

	if arguments.Get(1) != nil {
		return schema.Product{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Product{}, errors.New("error")
	} else {
		product := arguments.Get(0).(schema.Product)
		return product, nil
	}
}

func (repository *ProductRepositoryMock) FindById(ctx context.Context, productId string) (schema.Product, error) {

	arguments := repository.Mock.Called(ctx, productId)

	if arguments.Get(1) != nil {
		return schema.Product{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Product{}, errors.New("error")
	} else {
		product := arguments.Get(0).(schema.Product)
		return product, nil
	}
}

func (repository *ProductRepositoryMock) FindAll(ctx context.Context, skip int, limit int) ([]schema.Product, error) {

	arguments := repository.Mock.Called(ctx, skip, limit)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return []schema.Product{}, nil
	} else {
		return arguments.Get(0).([]schema.Product), nil
	}
}

func (repository *ProductRepositoryMock) FindAllWithSearch(ctx context.Context, search string, skip int, limit int) ([]schema.Product, error) {

	arguments := repository.Mock.Called(ctx, search, skip, limit)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return []schema.Product{}, nil
	} else {
		return arguments.Get(0).([]schema.Product), nil
	}
}

func (repository *ProductRepositoryMock) Update(ctx context.Context, product schema.Product) (schema.Product, error) {

	arguments := repository.Mock.Called(ctx, product)

	if arguments.Get(1) != nil {
		return schema.Product{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Product{}, errors.New("error")
	} else {
		product := arguments.Get(0).(schema.Product)
		return product, nil
	}
}

func (repository ProductRepositoryMock) PushImageIntoImages(ctx context.Context, productId string, images []schema.Image) ([]schema.Image, error) {

	arguments := repository.Mock.Called(ctx, productId, images)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return []schema.Image{}, nil
	} else {
		return arguments.Get(0).([]schema.Image), nil
	}
}

func (repository ProductRepositoryMock) PullImageFromImages(ctx context.Context, productId string, imageId string) (schema.Image, error) {

	arguments := repository.Mock.Called(ctx, productId, imageId)

	if arguments.Get(1) != nil {
		return schema.Image{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Image{}, errors.New("error")
	} else {
		return arguments.Get(0).(schema.Image), nil
	}
}

func (repository ProductRepositoryMock) Delete(ctx context.Context, productId string) error {

	arguments := repository.Mock.Called(ctx, productId)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}
}

func (repository ProductRepositoryMock) CountDocuments(ctx context.Context) (int, error) {

	arguments := repository.Mock.Called(ctx)

	if arguments.Get(1) != nil {
		return 0, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return 0, errors.New("error")
	} else {
		return arguments.Get(0).(int), nil
	}
}

func (repository ProductRepositoryMock) FindByMerchantId(ctx context.Context, merchantId string) ([]schema.Product, error) {

	arguments := repository.Mock.Called(ctx, merchantId)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return []schema.Product{}, nil
	} else {
		products := arguments.Get(0).([]schema.Product)
		return products, nil
	}

}

func (repository ProductRepositoryMock) FindByCategoryId(ctx context.Context, categoryId string) ([]schema.Product, error) {
	arguments := repository.Mock.Called(ctx, categoryId)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return []schema.Product{}, nil
	} else {
		products := arguments.Get(0).([]schema.Product)
		return products, nil
	}

}

func (repository ProductRepositoryMock) PullCategoryIdFromProduct(ctx context.Context, categoryId string) error {

	arguments := repository.Mock.Called(ctx, categoryId)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}

}

func (repository ProductRepositoryMock) UpdateQuantity(ctx context.Context, product schema.Product) error {

	arguments := repository.Mock.Called(ctx, product)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}
}
