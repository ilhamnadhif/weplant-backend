package service

import (
	"context"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type CustomerServiceImpl struct {
	CustomerRepository   repository.CustomerRepository
	ProductRepository    repository.ProductRepository
	CloudinaryRepository repository.CloudinaryRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository, productRepository repository.ProductRepository, cloudinaryRepository repository.CloudinaryRepository) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository:   customerRepository,
		ProductRepository:    productRepository,
		CloudinaryRepository: cloudinaryRepository,
	}
}

func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	res, err := service.CustomerRepository.Create(ctx, domain.Customer{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Email:     request.Email,
		Password:  helper.HashPassword(request.Password),
		UserName:  request.UserName,
		Phone:     request.Phone,
	})
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

func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId string) web.CustomerResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfErrorNotFound(err)

	return web.CustomerResponse{
		Id:        customer.Id.Hex(),
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
		Email:     customer.Email,
		UserName:  customer.UserName,
		Phone:     customer.Phone,
	}
}

func (service *CustomerServiceImpl) FindCartById(ctx context.Context, customerId string) web.CartResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfErrorNotFound(err)

	var totalPrice int
	var productsResponse []web.CartProductResponse

	for _, product := range customer.Carts {
		findProduct, err := service.ProductRepository.FindById(ctx, product.ProductId)
		helper.PanicIfError(err)
		subTotal := product.Quantity * findProduct.Price
		productsResponse = append(productsResponse, web.CartProductResponse{
			ProductId:   findProduct.Id.Hex(),
			Name:        findProduct.Name,
			Slug:        findProduct.Slug,
			Description: findProduct.Description,
			Price:       findProduct.Price,
			Quantity:    product.Quantity,
			SubTotal:    subTotal,
			MainImage: &web.ImageResponse{
				Id:       findProduct.MainImage.Id.Hex(),
				FileName: findProduct.MainImage.FileName,
				URL:      findProduct.MainImage.URL,
			},
		})
		totalPrice += subTotal
	}

	return web.CartResponse{
		CustomerId: customer.Id.Hex(),
		TotalPrice: totalPrice,
		Products:   productsResponse,
	}
}
func (service *CustomerServiceImpl) FindTransactionById(ctx context.Context, customerId string) web.TransactionResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfErrorNotFound(err)

	var transactionsResponse []web.TransactionDetailResponse

	for _, v := range customer.Transactions {
		var totalPrice int
		var productsResponse []web.TransactionProductResponse
		for _, p := range v.Products {
			product, err := service.ProductRepository.FindById(ctx, p.ProductId)
			helper.PanicIfError(err)

			subTotal := p.Price * p.Quantity

			totalPrice += subTotal

			productsResponse = append(productsResponse, web.TransactionProductResponse{
				ProductId:   product.Id.Hex(),
				Name:        product.Name,
				Slug:        product.Slug,
				Description: product.Description,
				Price:       p.Price,
				Quantity:    p.Quantity,
				SubTotal:    subTotal,
				MainImage: &web.ImageResponse{
					Id:       product.MainImage.Id.Hex(),
					FileName: product.MainImage.FileName,
					URL:      product.MainImage.URL,
				},
			})
		}

		transactionsResponse = append(transactionsResponse, web.TransactionDetailResponse{
			Id:         v.Id.Hex(),
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			Status:     v.Status,
			QRCode:     v.QRCode,
			TotalPrice: totalPrice,
			Products:   productsResponse,
			Address: &web.AddressResponse{
				Address:    v.Address.Address,
				City:       v.Address.City,
				Province:   v.Address.Province,
				Country:    v.Address.Country,
				PostalCode: v.Address.PostalCode,
			},
		})
	}

	return web.TransactionResponse{
		CustomerId:   customer.Id.Hex(),
		Transactions: transactionsResponse,
	}
}

func (service *CustomerServiceImpl) FindOrderById(ctx context.Context, customerId string) web.OrderResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfErrorNotFound(err)

	var productsResponse []web.OrderProductResponse
	for _, v := range customer.Orders {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)
		productsResponse = append(productsResponse, web.OrderProductResponse{
			Id:          v.Id.Hex(),
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			ProductId:   product.Id.Hex(),
			Name:        product.Name,
			Slug:        product.Slug,
			Description: product.Description,
			Price:       v.Price,
			Quantity:    v.Quantity,
			SubTotal:    v.Price * v.Quantity,
			MainImage: &web.ImageResponse{
				Id:       product.MainImage.Id.Hex(),
				FileName: product.MainImage.FileName,
				URL:      product.MainImage.URL,
			},
			Address: &web.AddressResponse{
				Address:    v.Address.Address,
				City:       v.Address.City,
				Province:   v.Address.Province,
				Country:    v.Address.Country,
				PostalCode: v.Address.PostalCode,
			},
		})
	}

	return web.OrderResponse{
		CustomerId: customer.Id.Hex(),
		Products:   productsResponse,
	}

}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerUpdateRequest {
	customer, err := service.CustomerRepository.FindById(ctx, request.Id)
	helper.PanicIfErrorNotFound(err)

	_, err = service.CustomerRepository.Update(ctx, domain.Customer{
		Id:        customer.Id,
		UpdatedAt: request.UpdatedAt,
		UserName:  request.UserName,
		Phone:     request.Phone,
	})
	helper.PanicIfError(err)
	return request
}

func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId string) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfErrorNotFound(err)

	err = service.CustomerRepository.Delete(ctx, customer.Id.Hex())
	helper.PanicIfError(err)
}
