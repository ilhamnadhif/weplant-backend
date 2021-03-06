package service

import (
	"context"
	"weplant-backend/model/web"
)

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) web.TokenResponse
	FindById(ctx context.Context, customerId string) web.CustomerResponse
	FindCartById(ctx context.Context, customerId string) web.CartResponse
	FindTransactionById(ctx context.Context, customerId string) web.TransactionResponse
	FindOrderById(ctx context.Context, customerId string) web.OrderResponse
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerUpdateRequest
	UpdateMainImage(ctx context.Context, request web.CustomerUpdateImageRequest) web.CustomerUpdateImageRequestResponse
	Delete(ctx context.Context, customerId string)
}
