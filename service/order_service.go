package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type OrderService interface {
	CheckoutFromCart(ctx context.Context, request web.OrderCreateRequest) snap.Response
	CallbackTransaction(ctx context.Context, response coreapi.TransactionStatusResponse) coreapi.TransactionStatusResponse
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

func (service *orderServiceImpl) CheckoutFromCart(ctx context.Context, request web.OrderCreateRequest) snap.Response {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfError(err)

	var totalPrice int
	var itemDetail []midtrans.ItemDetails

	for _, v := range customer.Carts {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)
		merchant, err := service.MerchantRepository.FindById(ctx, product.MerchantId)
		helper.PanicIfError(err)

		totalPrice += v.Quantity * product.Price

		helper.PanicIfError(err)

		itemDetail = append(itemDetail, midtrans.ItemDetails{
			ID:           product.Id.Hex(),
			Name:         product.Name,
			Price:        int64(product.Price),
			Qty:          int32(v.Quantity),
			MerchantName: merchant.Name,
		})
	}

	res, errorMidtrans := service.MidtransRepository.Checkout(snap.Request{
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
				Postcode:    request.Address.PostalCode,
				CountryCode: "IDN",
			},
			ShipAddr: &midtrans.CustomerAddress{
				FName:       customer.UserName,
				Phone:       customer.Phone,
				Address:     request.Address.Address,
				City:        request.Address.City,
				Postcode:    request.Address.PostalCode,
				CountryCode: "IDN",
			},
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		Metadata: struct {
			CustomerId string
			Address    string
			City       string
			Province   string
			Country    string
			PostalCode string
			Latitude   float64
			Longitude  float64
		}{
			CustomerId: customer.Id.Hex(),
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			Country:    request.Address.Country,
			PostalCode: request.Address.PostalCode,
			Latitude:   request.Address.Latitude,
			Longitude:  request.Address.Longitude,
		},
	})

	if errorMidtrans != nil {
		panic(errorMidtrans.Message)
	}

	return *res
}
func (service *orderServiceImpl) CallbackTransaction(ctx context.Context, request coreapi.TransactionStatusResponse) coreapi.TransactionStatusResponse {
	timeNow := helper.GetTimeNow()

	res, errCheck := service.MidtransRepository.CheckTransaction(request.OrderID)
	if errCheck != nil {
		panic(errCheck.Message)
	}

	fmt.Println(res)

	if !helper.CheckTransactionStatus(*res) {
		panic(errors.New("failed update order").Error())
	}

	orderRequest := res.Metadata.(map[string]interface{})
	fmt.Println("=========================================================================================================")
	fmt.Println(orderRequest)

	customer, err := service.CustomerRepository.FindById(ctx, orderRequest["CustomerId"].(string))
	helper.PanicIfError(err)

	var orderProduct []*domain.OrderProduct

	for _, v := range customer.Carts {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)

		orderProduct = append(orderProduct, &domain.OrderProduct{
			ProductId: product.Id.Hex(),
			Price:     product.Price,
			Quantity:  v.Quantity,
		})

		err = service.CustomerRepository.PullProductFromCart(ctx, customer.Id.Hex(), product.Id.Hex())
		helper.PanicIfError(err)

		merchant, err := service.MerchantRepository.FindById(ctx, product.MerchantId)
		helper.PanicIfError(err)

		_, err = service.MerchantRepository.Update(ctx, domain.Merchant{
			Id:        merchant.Id,
			UpdatedAt: timeNow,
			Balance:   merchant.Balance + (v.Quantity * product.Price),
		})

		_, err = service.ProductRepository.Update(ctx, domain.Product{
			Id:        product.Id,
			UpdatedAt: timeNow,
			Stock:     product.Stock - v.Quantity,
		})
		helper.PanicIfError(err)
	}

	err = service.CustomerRepository.CheckoutFromCart(ctx, customer.Id.Hex(), domain.Order{
		Id:        helper.ObjectIDFromHex(res.OrderID),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Products:  orderProduct,
		Address: &domain.Address{
			Address:    orderRequest["Address"].(string),
			City:       orderRequest["City"].(string),
			Province:   orderRequest["Province"].(string),
			Country:    orderRequest["Country"].(string),
			PostalCode: orderRequest["PostalCode"].(string),
			Latitude:   orderRequest["Latitude"].(float64),
			Longitude:  orderRequest["Longitude"].(float64),
		},
	})
	helper.PanicIfError(err)

	return *res
}
