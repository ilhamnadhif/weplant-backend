package service

import (
	"context"
	"weplant-backend/model/web"
)

type AuthService interface {
	LoginCustomer(ctx context.Context, request web.LoginRequest) web.TokenResponse
	LoginMerchant(ctx context.Context, request web.LoginRequest) web.TokenResponse
	LoginAdmin(ctx context.Context, request web.LoginRequest) web.TokenResponse
}
