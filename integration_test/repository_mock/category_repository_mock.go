package repository_mock

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"weplant-backend/model/schema"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) Create(ctx context.Context, category schema.Category) (schema.Category, error) {
	arguments := repository.Mock.Called(ctx, category)

	if arguments.Get(1) != nil {
		return schema.Category{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Category{}, errors.New("error")
	} else {
		category := arguments.Get(0).(schema.Category)
		return category, nil
	}
}

func (repository *CategoryRepositoryMock) FindById(ctx context.Context, categoryId string) (schema.Category, error) {
	arguments := repository.Mock.Called(ctx, categoryId)

	if arguments.Get(1) != nil {
		return schema.Category{}, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return schema.Category{}, errors.New("error")
	} else {
		category := arguments.Get(0).(schema.Category)
		return category, nil
	}
}

func (repository *CategoryRepositoryMock) FindAll(ctx context.Context) ([]schema.Category, error) {
	arguments := repository.Mock.Called(ctx)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return []schema.Category{}, nil
	} else {
		return arguments.Get(0).([]schema.Category), nil
	}

}

func (repository *CategoryRepositoryMock) Update(ctx context.Context, category schema.Category) (schema.Category, error) {
	panic("implement me")
}

func (repository *CategoryRepositoryMock) Delete(ctx context.Context, categoryId string) error {
	panic("implement me")
}
