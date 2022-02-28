package service

import (
	"context"
	"weplant-backend/model/web"
)

type MerchantService interface {
	Create(ctx context.Context, request web.MerchantCreateRequest) web.TokenResponse
	FindById(ctx context.Context, merchantId string) web.MerchantDetailResponse
	FindManageOrderById(ctx context.Context, merchantId string) web.ManageOrderResponse
	Update(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantUpdateRequest
	UpdateMainImage(ctx context.Context, request web.MerchantUpdateImageRequest) web.MerchantUpdateImageRequestResponse
	Delete(ctx context.Context, merchantId string)
}
