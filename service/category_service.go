package service

import (
	"context"
	"weplant-backend/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateRequestResponse
	FindById(ctx context.Context, categoryId string) web.CategoryDetailResponse
	FindAll(ctx context.Context) []web.CategorySimpleResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateRequest
	UpdateMainImage(ctx context.Context, request web.CategoryUpdateImageRequest) web.CategoryUpdateImageRequestResponse
	Delete(ctx context.Context, categoryId string)
}
