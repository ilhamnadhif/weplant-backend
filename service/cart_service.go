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

	product, errProduct := service.ProductRepository.FindById(ctx, request.ProductId)
	helper.PanicIfError(errProduct)

	errPush := service.CustomerRepository.PushProductTCart(ctx, customer.Id.Hex(), domain.CartProduct{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		ProductId: product.Id.Hex(),
		Quantity:  request.Quantity,
	}, product.Price)
	helper.PanicIfError(errPush)
	return request
}

func (service *cartServiceImpl) UpdateProductQuantity(ctx context.Context, request web.CartProductUpdateRequest) web.CartProductUpdateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfError(err)

	product, errProduct := service.ProductRepository.FindById(ctx, request.ProductId)
	helper.PanicIfError(errProduct)

	errUpdate := service.CustomerRepository.UpdateProductQuantity(ctx, customer.Id.Hex(), domain.CartProduct{
		UpdatedAt: request.UpdatedAt,
		ProductId: product.Id.Hex(),
		Quantity:  request.Quantity,
	}, request.Quantity*product.Price)
	helper.PanicIfError(errUpdate)

	return request
}

func (service *cartServiceImpl) PullProductFromCart(ctx context.Context, customerId string, productId string) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	product, errProduct := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfError(errProduct)

	for _, prdct := range customer.Cart.Products {
		if prdct.ProductId == product.Id.Hex() {
			service.CustomerRepository.PullProductFromCart(ctx, customer.Id.Hex(), product.Id.Hex(), -prdct.Quantity*product.Price)
		}
	}
}
