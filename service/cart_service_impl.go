package service

import (
	"context"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type CartServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	ProductRepository  repository.ProductRepository
}

func NewCartService(customerRepository repository.CustomerRepository, productRepository repository.ProductRepository) CartService {
	return &CartServiceImpl{
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
	}
}

func (service *CartServiceImpl) PushProductToCart(ctx context.Context, request web.CartProductCreateRequest) web.CartProductCreateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfErrorNotFound(err)

	product, err := service.ProductRepository.FindById(ctx, request.ProductId)
	helper.PanicIfError(err)

	err = service.CustomerRepository.PushProductToCart(ctx, customer.Id.Hex(), schema.CartProduct{
		ProductId: product.Id.Hex(),
		Quantity:  request.Quantity,
	})
	helper.PanicIfError(err)
	return request
}

func (service *CartServiceImpl) UpdateProductQuantity(ctx context.Context, request web.CartProductUpdateRequest) web.CartProductUpdateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfErrorNotFound(err)

	product, err := service.ProductRepository.FindById(ctx, request.ProductId)
	helper.PanicIfErrorNotFound(err)

	err = service.CustomerRepository.UpdateProductQuantity(ctx, customer.Id.Hex(), schema.CartProduct{
		ProductId: product.Id.Hex(),
		Quantity:  request.Quantity,
	})
	helper.PanicIfError(err)

	return request
}

func (service *CartServiceImpl) PullProductFromCart(ctx context.Context, customerId string, productId string) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfErrorNotFound(err)

	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfErrorNotFound(err)

	for _, v := range customer.Carts {
		if v.ProductId == product.Id.Hex() {
			err = service.CustomerRepository.PullProductFromCart(ctx, customer.Id.Hex(), product.Id.Hex())
			helper.PanicIfError(err)
		}
	}
}
