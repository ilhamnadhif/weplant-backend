package service

import (
	"context"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type OrderService interface {
	CheckoutFromCart(ctx context.Context, request web.OrderProductCreateRequest) *snap.Response
}

type orderServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	ProductRepository  repository.ProductRepository
	MidtransRepository repository.MidtransRepository
	MerchantRepository repository.MerchantRepository
}

func NewOrderService(customerRepository repository.CustomerRepository, productRepository repository.ProductRepository, midtransRepository repository.MidtransRepository, merchantRepository repository.MerchantRepository) OrderService {
	return &orderServiceImpl{
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
		MidtransRepository: midtransRepository,
		MerchantRepository: merchantRepository,
	}
}

func (service *orderServiceImpl) CheckoutFromCart(ctx context.Context, request web.OrderProductCreateRequest) *snap.Response {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfError(err)

	var totalPrice int
	var itemDetail []midtrans.ItemDetails

	for _, v := range customer.Cart.Products {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)
		merchant, err := service.MerchantRepository.FindById(ctx, product.MerchantId)
		helper.PanicIfError(err)

		totalPrice += v.Quantity * product.Price

		err = service.CustomerRepository.CheckoutFromCart(ctx, customer.Id.Hex(), domain.OrderProduct{
			Id:        primitive.NewObjectID(),
			CreatedAt: request.CreatedAt,
			UpdatedAt: request.UpdatedAt,
			ProductId: product.Id.Hex(),
			Price:     product.Price,
			Quantity:  v.Quantity,
			HasDone:   helper.ReturnPointerBool(false),
			Address: &domain.Address{
				Address:    request.Address.Address,
				City:       request.Address.City,
				Province:   request.Address.Province,
				Country:    request.Address.Country,
				PostalCode: request.Address.PostalCode,
				Latitude:   request.Address.Latitude,
				Longitude:  request.Address.Longitude,
			},
		})
		helper.PanicIfError(err)

		_, err = service.ProductRepository.Update(ctx, domain.Product{
			UpdatedAt: request.UpdatedAt,
			Stock:     product.Stock - v.Quantity,
		})
		helper.PanicIfError(err)

		err = service.CustomerRepository.PullProductFromCart(ctx, customer.Id.Hex(), v.ProductId)
		helper.PanicIfError(err)

		itemDetail = append(itemDetail, midtrans.ItemDetails{
			ID:           product.Id.Hex(),
			Name:         product.Name,
			Price:        int64(product.Price),
			Qty:          int32(v.Quantity),
			MerchantName: merchant.Name,
		})
	}

	res, errorMidtrans := service.MidtransRepository.SendBalance(snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  primitive.NewObjectID().Hex(),
			GrossAmt: int64(totalPrice),
		},
		Items: &itemDetail,
		CustomerDetail: &midtrans.CustomerDetails{
			FName: customer.UserName,
			Email: customer.Email,
			Phone: customer.Phone,
			BillAddr: &midtrans.CustomerAddress{
				FName:       customer.UserName,
				Phone:       customer.Phone,
				Address:     request.Address.Address,
				City:        request.Address.City,
				Postcode:    strconv.Itoa(request.Address.PostalCode),
				CountryCode: "IDN",
			},
			ShipAddr: &midtrans.CustomerAddress{
				FName:       customer.UserName,
				Phone:       customer.Phone,
				Address:     request.Address.Address,
				City:        request.Address.City,
				Postcode:    strconv.Itoa(request.Address.PostalCode),
				CountryCode: "IDN",
			},
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		UserId: customer.Id.Hex(),
	})

	if errorMidtrans != nil {
		panic(errorMidtrans.Message)
	}

	return res
}
