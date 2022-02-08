package service

import (
	"context"
	"weplant-backend/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	FindById(ctx context.Context, categoryId string) web.CategoryResponseWithProduct
	FindAll(ctx context.Context) []web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateRequest
	UpdateMainImage(ctx context.Context, request web.CategoryUpdateImageRequest) web.CategoryUpdateImageRequest
	Delete(ctx context.Context, categoryId string)
}
