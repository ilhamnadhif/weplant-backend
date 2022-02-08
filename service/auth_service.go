package service

import (
	"context"
	"weplant-backend/model/web"
)

type AuthService interface {
	LoginCustomer(ctx context.Context, request web.LoginRequest) web.LoginResponse
	LoginMerchant(ctx context.Context, request web.LoginRequest) web.LoginResponse
	LoginAdmin(ctx context.Context, request web.LoginRequest) web.LoginResponse
}
