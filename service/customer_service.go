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
	FindTransactionById(ctx context.Context, customerId string) web.TransactionResponse
	FindOrderById(ctx context.Context, customerId string) web.OrderResponse
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
		subTotal := product.Quantity * findProduct.Price
		products = append(products, &web.CartProductResponse{
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
		Products:   products,
	}
}
func (service *customerServiceImpl) FindTransactionById(ctx context.Context, customerId string) web.TransactionResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	var transactionResponse []*web.TransactionDetailResponse

	for _, v := range customer.Transactions {

		var totalPrice int

		var productResponse []*web.TransactionProductResponse

		for _, p := range v.Products {
			product, err := service.ProductRepository.FindById(ctx, p.ProductId)
			helper.PanicIfError(err)

			subtotal := p.Price * p.Quantity

			totalPrice += subtotal

			productResponse = append(productResponse, &web.TransactionProductResponse{
				ProductId:   product.Id.Hex(),
				Name:        product.Name,
				Slug:        product.Slug,
				Description: product.Description,
				Price:       p.Price,
				Quantity:    p.Quantity,
				SubTotal:    subtotal,
				MainImage: &web.ImageResponse{
					Id:       product.MainImage.Id.Hex(),
					FileName: product.MainImage.FileName,
					URL:      product.MainImage.URL,
				},
			})
		}

		transactionResponse = append(transactionResponse, &web.TransactionDetailResponse{
			Id:         v.Id.Hex(),
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			Status:     v.Status,
			QRCode:     v.QRCode,
			TotalPrice: totalPrice,
			Products:   productResponse,
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
		Transactions: transactionResponse,
	}
}

func (service *customerServiceImpl) FindOrderById(ctx context.Context, customerId string) web.OrderResponse {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfError(err)

	var productResponse []*web.OrderProductResponse
	for _, v := range customer.Orders {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)
		productResponse = append(productResponse, &web.OrderProductResponse{
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
		Products:   productResponse,
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
