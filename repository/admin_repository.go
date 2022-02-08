package repository

import (
	"context"
	"weplant-backend/model/domain"
)

type AdminRepository interface {
	Create(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	FindById(ctx context.Context, adminId string) (domain.Admin, error)
	FindByEmail(ctx context.Context, email string) (domain.Admin, error)
	FindAll(ctx context.Context) ([]domain.Admin, error)
	Update(ctx context.Context, admin domain.Admin) (domain.Admin, error)
}
