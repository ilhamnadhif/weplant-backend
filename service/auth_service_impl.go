package service

import (
	"context"
	"errors"
	"weplant-backend/helper"
	"weplant-backend/model/web"
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

func (service *AuthServiceImpl) LoginCustomer(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	customer, err := service.CustomerRepository.FindByEmail(ctx, request.Email)
	helper.PanicIfError(err)
	if !helper.CheckPasswordHash(request.Password, customer.Password) {
		panic(errors.New("password not match").Error())
	}
	return web.LoginResponse{
		Id:   customer.Id.Hex(),
		Role: "customer",
	}
}

func (service *AuthServiceImpl) LoginMerchant(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	merchant, err := service.MerchantRepository.FindByEmail(ctx, request.Email)
	helper.PanicIfError(err)
	if !helper.CheckPasswordHash(request.Password, merchant.Password) {
		panic(errors.New("password not match").Error())
	}

	return web.LoginResponse{
		Id:   merchant.Id.Hex(),
		Role: "merchant",
	}
}

func (service *AuthServiceImpl) LoginAdmin(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	admin, err := service.AdminRepository.FindByEmail(ctx, request.Email)
	helper.PanicIfError(err)
	if !helper.CheckPasswordHash(request.Password, admin.Password) {
		panic(errors.New("password not match").Error())
	}
	return web.LoginResponse{
		Id:   admin.Id.Hex(),
		Role: "admin",
	}
}
