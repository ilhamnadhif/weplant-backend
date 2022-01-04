package service

import (
	"context"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerCreateRequest
	FindById(ctx context.Context, customerId string) web.CustomerResponse
	FindCartById(ctx context.Context, customerId string) web.CartResponse
	FindOrderById(ctx context.Context, customerId string) web.CustomerOrdersResponse
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerUpdateRequest
	Delete(ctx context.Context, customerId string)
}

type customerServiceImpl struct {
	CustomerRepository   repository.CustomerRepository
	ProductRepository    repository.ProductRepository
	CloudinaryRepository repository.CloudinaryRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository, productRepository repository.ProductRepository, cloudinaryRepository repository.CloudinaryRepository) CustomerService {
	return &customerServiceImpl{
		CustomerRepository:   customerRepository,
		ProductRepository:    productRepository,
		CloudinaryRepository: cloudinaryRepository,
	}
}

func (service *customerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerCreateRequest {
	_, err := service.CustomerRepository.Create(ctx, domain.Customer{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Email:     request.Email,
		Password:  helper.HashPassword(request.Password),
		UserName:  request.UserName,
		Phone:     request.Phone,
	})
	helper.PanicIfError(err)
	return request
}

func (service *customerServiceImpl) FindById(ctx context.Context, customerId string) web.CustomerResponse {
	res, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	return web.CustomerResponse{
		Id:        res.Id.Hex(),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Email:     res.Email,
		UserName:  res.UserName,
		Phone:     res.Phone,
	}
}

func (service *customerServiceImpl) FindCartById(ctx context.Context, customerId string) web.CartResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	var totalPrice int

	var products []*web.CartProductResponse
	for _, product := range customer.Carts {
		findProduct, err := service.ProductRepository.FindById(ctx, product.ProductId)
		helper.PanicIfError(err)
		products = append(products, &web.CartProductResponse{
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			ProductId:   findProduct.Id.Hex(),
			Name:        findProduct.Name,
			Description: findProduct.Description,
			Price:       findProduct.Price,
			MainImage: &web.ImageResponse{
				Id:       findProduct.MainImage.Id.Hex(),
				FileName: findProduct.MainImage.FileName,
				URL:      findProduct.MainImage.URL,
			},
			Quantity: product.Quantity,
		})
		totalPrice += product.Quantity * findProduct.Price
	}

	return web.CartResponse{
		CustomerId: customer.Id.Hex(),
		Total:      totalPrice,
		Products:   products,
	}
}

func (service *customerServiceImpl) FindOrderById(ctx context.Context, customerId string) web.CustomerOrdersResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	var orders []*web.OrderResponse
	for _, order := range customer.Orders {
		var productsResponse []*web.OrderProductResponse
		for _, v := range order.Products {
			product, err := service.ProductRepository.FindById(ctx, v.ProductId)
			helper.PanicIfError(err)
			productsResponse = append(productsResponse, &web.OrderProductResponse{
				ProductId:   product.Id.Hex(),
				Name:        product.Name,
				Description: product.Description,
				Price:       v.Price,
				MainImage: &web.ImageResponse{
					Id:       product.MainImage.Id.Hex(),
					FileName: product.MainImage.FileName,
					URL:      product.MainImage.URL,
				},
				Quantity: v.Quantity,
			})
		}

		orders = append(orders, &web.OrderResponse{
			Id:        order.Id.Hex(),
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			Products:  productsResponse,
			Address: &web.AddressResponse{
				Address:    order.Address.Address,
				City:       order.Address.City,
				Province:   order.Address.Province,
				Country:    order.Address.Country,
				PostalCode: order.Address.PostalCode,
				Latitude:   order.Address.Latitude,
				Longitude:  order.Address.Longitude,
			},
		})
	}

	return web.CustomerOrdersResponse{
		CustomerId: customer.Id.Hex(),
		Orders:     orders,
	}
}

func (service *customerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerUpdateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	_, err = service.CustomerRepository.Update(ctx, domain.Customer{
		Id:        customer.Id,
		UpdatedAt: request.UpdatedAt,
		UserName:  request.UserName,
		Phone:     request.Phone,
	})
	helper.PanicIfError(err)
	return request
}

func (service *customerServiceImpl) Delete(ctx context.Context, customerId string) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	err = service.CustomerRepository.Delete(ctx, customer.Id.Hex())
	helper.PanicIfError(err)
}
