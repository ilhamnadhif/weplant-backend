package repository

import (
	"context"
	"weplant-backend/model/schema"
)

type ProductRepository interface {
	Create(ctx context.Context, product schema.Product) (schema.Product, error)
	FindById(ctx context.Context, productId string) (schema.Product, error)
	FindAll(ctx context.Context, skip int, limit int) ([]schema.Product, error)
	FindAllWithSearch(ctx context.Context, search string, skip int, limit int) ([]schema.Product, error)
	Update(ctx context.Context, product schema.Product) (schema.Product, error)
	PushImageIntoImages(ctx context.Context, productId string, images []schema.Image) ([]schema.Image, error)
	PullImageFromImages(ctx context.Context, productId string, imageId string) (schema.Image, error)
	Delete(ctx context.Context, productId string) error
	CountDocuments(ctx context.Context) (int, error)

	// merchant
	FindByMerchantId(ctx context.Context, merchantId string) ([]schema.Product, error)

	// category
	FindByCategoryId(ctx context.Context, categoryId string) ([]schema.Product, error)
	PullCategoryIdFromProduct(ctx context.Context, categoryId string) error

	// transaction
	UpdateQuantity(ctx context.Context, product schema.Product) error
}
