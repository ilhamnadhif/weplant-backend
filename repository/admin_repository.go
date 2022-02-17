package repository

import (
	"context"
	"weplant-backend/model/schema"
)

type AdminRepository interface {
	Create(ctx context.Context, admin schema.Admin) (schema.Admin, error)
	FindById(ctx context.Context, adminId string) (schema.Admin, error)
	FindByEmail(ctx context.Context, email string) (schema.Admin, error)
	FindAll(ctx context.Context) ([]schema.Admin, error)
	Update(ctx context.Context, admin schema.Admin) (schema.Admin, error)
}
