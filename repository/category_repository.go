package repository

import (
	"context"
	"weplant-backend/model/schema"
)

type CategoryRepository interface {
	Create(ctx context.Context, category schema.Category) (schema.Category, error)
	FindById(ctx context.Context, categoryId string) (schema.Category, error)
	FindAll(ctx context.Context) ([]schema.Category, error)
	Update(ctx context.Context, category schema.Category) (schema.Category, error)
	Delete(ctx context.Context, categoryId string) error
}
