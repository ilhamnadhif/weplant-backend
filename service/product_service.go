package service

import (
	"context"
	"weplant-backend/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductCreateRequestResponse
	FindById(ctx context.Context, productId string) web.ProductDetailResponse
	FindAll(ctx context.Context, page int, perPage int) web.ProductFindAllResponse
	FindAllWithSearch(ctx context.Context, search string, page int, perPage int) web.ProductFindAllResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductUpdateRequest
	UpdateMainImage(ctx context.Context, request web.ProductUpdateImageRequest) web.ProductUpdateImageRequestResponse
	PushImageIntoImages(ctx context.Context, productId string, request []web.ImageCreateRequest) []web.ImageCreateRequest
	PullImageFromImages(ctx context.Context, productId string, imageId string)
	Delete(ctx context.Context, productId string)
}
