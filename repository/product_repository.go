package repository

import (
	"context"
	"weplant-backend/model/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	FindById(ctx context.Context, productId string) (domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindAllWithSearch(ctx context.Context, search string) ([]domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	PushImageIntoImages(ctx context.Context, productId string, images []domain.Image) ([]domain.Image, error)
	PullImageFromImages(ctx context.Context, productId string, imageId string) (domain.Image, error)
	Delete(ctx context.Context, productId string) error

	// merchant
	FindByMerchantId(ctx context.Context, merchantId string) ([]domain.Product, error)

	// category
	FindByCategoryId(ctx context.Context, categoryId string) ([]domain.Product, error)
	PullCategoryIdFromProduct(ctx context.Context, categoryId string) error

	// transaction
	UpdateQuantity(ctx context.Context, product domain.Product) error
}
