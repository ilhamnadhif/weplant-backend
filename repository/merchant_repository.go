package repository

import (
	"context"
	"weplant-backend/model/domain"
)

type MerchantRepository interface {
	Create(ctx context.Context, merchant domain.Merchant) (domain.Merchant, error)
	FindById(ctx context.Context, merchantId string) (domain.Merchant, error)
	FindByEmail(ctx context.Context, email string) (domain.Merchant, error)
	FindBySlug(ctx context.Context, slug string) (domain.Merchant, error)
	Update(ctx context.Context, merchant domain.Merchant) (domain.Merchant, error)
	//UpdateBalance(ctx context.Context, merchant domain.Merchant) error
	Delete(ctx context.Context, merchantId string) error

	// Manage Order
	PushProductToManageOrders(ctx context.Context, merchantId string, product domain.ManageOrderProduct) error
}
