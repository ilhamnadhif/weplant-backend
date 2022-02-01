package service

import (
	"context"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type CartService interface {
	PushProductToCart(ctx context.Context, request web.CartProductCreateRequest) web.CartProductCreateRequest
	UpdateProductQuantity(ctx context.Context, request web.CartProductUpdateRequest) web.CartProductUpdateRequest
	PullProductFromCart(ctx context.Context, customerId string, productId string)
}

type cartServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	ProductRepository  repository.ProductRepository
}

func NewCartService(customerRepository repository.CustomerRepository, productRepository repository.ProductRepository) CartService {
	return &cartServiceImpl{
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
	}
}

func (service *cartServiceImpl) PushProductToCart(ctx context.Context, request web.CartProductCreateRequest) web.CartProductCreateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfError(err)

	product, err := service.ProductRepository.FindById(ctx, request.ProductId)
	helper.PanicIfError(err)

	err = service.CustomerRepository.PushProductToCart(ctx, customer.Id.Hex(), domain.CartProduct{
		ProductId: product.Id.Hex(),
		Quantity:  request.Quantity,
	})
	helper.PanicIfError(err)
	return request
}

func (service *cartServiceImpl) UpdateProductQuantity(ctx context.Context, request web.CartProductUpdateRequest) web.CartProductUpdateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfError(err)

	product, err := service.ProductRepository.FindById(ctx, request.ProductId)
	helper.PanicIfError(err)

	errUpdate := service.CustomerRepository.UpdateProductQuantity(ctx, customer.Id.Hex(), domain.CartProduct{
		ProductId: product.Id.Hex(),
		Quantity:  request.Quantity,
	})
	helper.PanicIfError(errUpdate)

	return request
}

func (service *cartServiceImpl) PullProductFromCart(ctx context.Context, customerId string, productId string) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfError(err)

	for _, prdct := range customer.Carts {
		if prdct.ProductId == product.Id.Hex() {
			service.CustomerRepository.PullProductFromCart(ctx, customer.Id.Hex(), product.Id.Hex())
		}
	}
}
