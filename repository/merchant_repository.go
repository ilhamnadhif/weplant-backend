package repository

import (
	"context"
	"weplant-backend/model/schema"
)

type MerchantRepository interface {
	Create(ctx context.Context, merchant schema.Merchant) (schema.Merchant, error)
	FindById(ctx context.Context, merchantId string) (schema.Merchant, error)
	FindByEmail(ctx context.Context, email string) (schema.Merchant, error)
	FindBySlug(ctx context.Context, slug string) (schema.Merchant, error)
	Update(ctx context.Context, merchant schema.Merchant) (schema.Merchant, error)
	//UpdateBalance(ctx context.Context, merchant schema.Merchant) error
	Delete(ctx context.Context, merchantId string) error

	// Manage Order
	PushProductToManageOrders(ctx context.Context, merchantId string, product schema.ManageOrderProduct) error
}
