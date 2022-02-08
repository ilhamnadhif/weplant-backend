package service

import (
	"context"
	"weplant-backend/model/web"
)

type CartService interface {
	PushProductToCart(ctx context.Context, request web.CartProductCreateRequest) web.CartProductCreateRequest
	UpdateProductQuantity(ctx context.Context, request web.CartProductUpdateRequest) web.CartProductUpdateRequest
	PullProductFromCart(ctx context.Context, customerId string, productId string)
}
