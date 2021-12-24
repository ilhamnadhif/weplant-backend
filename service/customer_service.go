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
	for _, product := range customer.Cart.Products {
		findProduct, errPrdct := service.ProductRepository.FindById(ctx, product.ProductId)
		helper.PanicIfError(errPrdct)
		url, errUrl := service.CloudinaryRepository.GetImage(ctx, findProduct.MainImage.FileName)
		helper.PanicIfError(errUrl)
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
				URL:      url,
			},
			Quantity: product.Quantity,
		})
		totalPrice += product.Quantity * findProduct.Price
	}

	return web.CartResponse{
		CustomerId: customer.Id.Hex(),
		Total: totalPrice,
		Products: products,
	}
}

func (service *customerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerUpdateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	_, errUpdate := service.CustomerRepository.Update(ctx, domain.Customer{
		Id:        customer.Id,
		UpdatedAt: request.UpdatedAt,
		UserName:  request.UserName,
		Phone:     request.Phone,
	})
	helper.PanicIfError(errUpdate)
	return request
}

func (service *customerServiceImpl) Delete(ctx context.Context, customerId string) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	errDelete := service.CustomerRepository.Delete(ctx, customer.Id.Hex())
	helper.PanicIfError(errDelete)
}
