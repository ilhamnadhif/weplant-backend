package service

import (
	"context"
	"weplant-backend/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateRequestResponse
	FindById(ctx context.Context, categoryId string) web.CategoryDetailResponse
	FindAll(ctx context.Context) []web.CategorySimpleResponse
}
