package service

import (
	"context"
	"weplant-backend/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductSimpleResponse
	FindById(ctx context.Context, productId string) web.ProductDetailResponse
	FindAll(ctx context.Context) []web.ProductSimpleResponse
	FindAllWithSearch(ctx context.Context, search string) []web.ProductSimpleResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductUpdateRequest
	UpdateMainImage(ctx context.Context, request web.ProductUpdateImageRequest) web.ProductUpdateImageRequest
	PushImageIntoImages(ctx context.Context, productId string, request []web.ImageCreateRequest) []web.ImageCreateRequest
	PullImageFromImages(ctx context.Context, productId string, imageId string)
	Delete(ctx context.Context, productId string)
}
