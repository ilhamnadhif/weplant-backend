package service

import (
	"context"
	"errors"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/pkg"
	"weplant-backend/repository"
)

type AuthServiceImpl struct {
	MerchantRepository repository.MerchantRepository
	CustomerRepository repository.CustomerRepository
	AdminRepository    repository.AdminRepository
}

func NewAuthService(merchantRepository repository.MerchantRepository, customerRepository repository.CustomerRepository, adminRepository repository.AdminRepository) AuthService {
	return &AuthServiceImpl{
		MerchantRepository: merchantRepository,
		CustomerRepository: customerRepository,
		AdminRepository:    adminRepository,
	}
}

func (service *AuthServiceImpl) LoginCustomer(ctx context.Context, request web.LoginRequest) web.TokenResponse {
	customer, err := service.CustomerRepository.FindByEmail(ctx, request.Email)
	helper.PanicIfError(err)
	if !pkg.CheckPasswordHash(request.Password, customer.Password) {
		panic(errors.New("password not match").Error())
	}
	token := pkg.GenerateToken(web.JWTPayload{
		Id:   customer.Id.Hex(),
		Role: "customer",
	})
	return web.TokenResponse{
		Id:    customer.Id.Hex(),
		Role:  "customer",
		Token: token,
	}
}

func (service *AuthServiceImpl) LoginMerchant(ctx context.Context, request web.LoginRequest) web.TokenResponse {
	merchant, err := service.MerchantRepository.FindByEmail(ctx, request.Email)
	helper.PanicIfError(err)
	if !pkg.CheckPasswordHash(request.Password, merchant.Password) {
		panic(errors.New("password not match").Error())
	}

	token := pkg.GenerateToken(web.JWTPayload{
		Id:   merchant.Id.Hex(),
		Role: "merchant",
	})
	return web.TokenResponse{
		Id:    merchant.Id.Hex(),
		Role:  "merchant",
		Token: token,
	}
}

func (service *AuthServiceImpl) LoginAdmin(ctx context.Context, request web.LoginRequest) web.TokenResponse {
	admin, err := service.AdminRepository.FindByEmail(ctx, request.Email)
	helper.PanicIfError(err)
	if !pkg.CheckPasswordHash(request.Password, admin.Password) {
		panic(errors.New("password not match").Error())
	}

	token := pkg.GenerateToken(web.JWTPayload{
		Id:   admin.Id.Hex(),
		Role: "admin",
	})
	return web.TokenResponse{
		Id:    admin.Id.Hex(),
		Role:  "admin",
		Token: token,
	}
}
