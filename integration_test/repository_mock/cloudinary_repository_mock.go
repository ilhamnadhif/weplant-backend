package repository_mock

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
)

type CloudinaryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CloudinaryRepositoryMock) UploadImage(ctx context.Context, filename string, image interface{}) (string, error) {
	arguments := repository.Mock.Called(ctx, filename, image)

	if arguments.Get(1) != nil {
		return "", arguments.Get(1).(error)
	}

	if arguments.Get(0) == nil {
		return "", errors.New("error")
	} else {
		return arguments.Get(0).(string), nil
	}
}

func (repository *CloudinaryRepositoryMock) DeleteImage(ctx context.Context, filename string) error {
	arguments := repository.Mock.Called(ctx, filename)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	} else {
		return nil
	}
}
