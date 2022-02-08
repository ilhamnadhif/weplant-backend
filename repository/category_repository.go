package repository

import (
	"context"
	"weplant-backend/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category domain.Category) (domain.Category, error)
	FindById(ctx context.Context, categoryId string) (domain.Category, error)
	FindAll(ctx context.Context) ([]domain.Category, error)
	Update(ctx context.Context, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, categoryId string) error
}
