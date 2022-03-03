package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type TransactionServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	ProductRepository  repository.ProductRepository
	MidtransRepository repository.MidtransRepository
	MerchantRepository repository.MerchantRepository
}

func NewTransactionService(customerRepository repository.CustomerRepository, productRepository repository.ProductRepository, midtransRepository repository.MidtransRepository, merchantRepository repository.MerchantRepository) TransactionService {
	return &TransactionServiceImpl{
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
		MidtransRepository: midtransRepository,
		MerchantRepository: merchantRepository,
	}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, request web.TransactionCreateRequest) web.TransactionCreateRequestResponse {
	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerId)
	helper.PanicIfErrorNotFound(err)

	var productDetailMidtrans []midtrans.ItemDetails
	var productDetailTransaction []schema.TransactionProduct

	var totalPrice int64

	for _, v := range customer.Carts {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)

		if v.Quantity > product.Stock {
			panic(fmt.Sprintf("barang %s yang anda beli harus kurang dari %d, dari stock yang tersedia", product.Name, product.Stock))
		} else if v.Quantity < 1 {
			panic(fmt.Sprintf("barang %s yang anda beli tidak boleh kurang dari 1", product.Name))
		}

		merchant, err := service.MerchantRepository.FindById(ctx, product.MerchantId)
		helper.PanicIfError(err)

		totalPrice += int64(product.Price * v.Quantity)

		productDetailMidtrans = append(productDetailMidtrans, midtrans.ItemDetails{
			ID:           product.Id.Hex(),
			Name:         product.Name,
			Price:        int64(product.Price),
			Qty:          int32(v.Quantity),
			MerchantName: merchant.Name,
		})

		productDetailTransaction = append(productDetailTransaction, schema.TransactionProduct{
			ProductId: product.Id.Hex(),
			Price:     product.Price,
			Quantity:  v.Quantity,
		})

		err = service.CustomerRepository.PullProductFromCart(ctx, customer.Id.Hex(), product.Id.Hex())
		helper.PanicIfError(err)
	}

	resMidtrans, errMidtrans := service.MidtransRepository.CreateTransaction(coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  primitive.NewObjectID().Hex(),
			GrossAmt: totalPrice,
		},
		Items: &productDetailMidtrans,
		CustomerDetails: &midtrans.CustomerDetails{
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
		CustomField1: helper.ReturnPointerString(customer.Id.Hex()),
	})
	if errMidtrans != nil {
		panic(errMidtrans.GetMessage())
	}

	var actionsCreateRequest []schema.TransactionAction
	var actionsResponse []web.TransactionActionResponse
	for _, action := range resMidtrans.Actions {
		if strings.ToLower(action.Name) == "generate-qr-code" {
			actionsCreateRequest = append(actionsCreateRequest, schema.TransactionAction{
				Name:   action.Name,
				Method: action.Method,
				URL:    action.URL,
			})
			actionsResponse = append(actionsResponse, web.TransactionActionResponse{
				Name:   action.Name,
				Method: action.Method,
				URL:    action.URL,
			})
		} else if strings.ToLower(action.Name) == "deeplink-redirect" {
			actionsCreateRequest = append(actionsCreateRequest, schema.TransactionAction{
				Name:   action.Name,
				Method: action.Method,
				URL:    action.URL,
			})
			actionsResponse = append(actionsResponse, web.TransactionActionResponse{
				Name:   action.Name,
				Method: action.Method,
				URL:    action.URL,
			})
		}
	}

	err = service.CustomerRepository.CreateTransaction(ctx, customer.Id.Hex(), schema.Transaction{
		Id:          helper.ObjectIDFromHex(resMidtrans.OrderID),
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		PaymentType: resMidtrans.PaymentType,
		Status:      resMidtrans.TransactionStatus,
		Actions:     actionsCreateRequest,
		Products:    productDetailTransaction,
		Address: &schema.Address{
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			PostalCode: request.Address.PostalCode,
		},
	})
	helper.PanicIfError(err)

	return web.TransactionCreateRequestResponse{
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		PaymentType: resMidtrans.PaymentType,
		Status:      resMidtrans.TransactionStatus,
		Actions:     actionsResponse,
		TotalPrice:  int(totalPrice),
		Address:     *request.Address,
	}
}

func (service *TransactionServiceImpl) Cancel(ctx context.Context, customerId string, transactionId string) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	helper.PanicIfErrorNotFound(err)

	var found bool

	for _, transaction := range customer.Transactions {
		if transaction.Id.Hex() == transactionId {
			found = true
			break
		}
	}

	if !found {
		helper.PanicIfErrorNotFound(errors.New(fmt.Sprintf("transaction id %s not found in customer id %s ", customerId, transactionId)))
	}

	_, errMidtrans := service.MidtransRepository.CancelTransaction(transactionId)
	if errMidtrans != nil {
		panic(errMidtrans.GetMessage())
	}
}

func (service *TransactionServiceImpl) Callback(ctx context.Context, request coreapi.TransactionStatusResponse) {
	timeNow := helper.GetTimeNow()

	res, errMidtrans := service.MidtransRepository.CheckTransaction(request.OrderID)
	if errMidtrans != nil {
		panic(errMidtrans.GetMessage())
	}

	customer, err := service.CustomerRepository.FindById(ctx, res.CustomField1)
	helper.PanicIfError(err)

	switch helper.CheckTransactionStatus(*res) {
	case "success":
		for _, v := range customer.Transactions {
			if v.Id.Hex() == res.OrderID {
				for _, p := range v.Products {
					product, err := service.ProductRepository.FindById(ctx, p.ProductId)
					helper.PanicIfError(err)
					err = service.CustomerRepository.CreateOrder(ctx, customer.Id.Hex(), schema.OrderProduct{
						Id:        primitive.NewObjectID(),
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
						ProductId: product.Id.Hex(),
						Price:     p.Price,
						Quantity:  p.Quantity,
						Address: &schema.Address{
							Address:    v.Address.Address,
							City:       v.Address.City,
							Province:   v.Address.Province,
							PostalCode: v.Address.PostalCode,
						},
					})
					helper.PanicIfError(err)
					err = service.MerchantRepository.PushProductToManageOrders(ctx, product.MerchantId, schema.ManageOrderProduct{
						Id:        primitive.NewObjectID(),
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
						ProductId: product.Id.Hex(),
						Price:     p.Price,
						Quantity:  p.Quantity,
						Address: &schema.Address{
							Address:    v.Address.Address,
							City:       v.Address.City,
							Province:   v.Address.Province,
							PostalCode: v.Address.PostalCode,
						},
					})
					helper.PanicIfError(err)
					err = service.ProductRepository.UpdateQuantity(ctx, schema.Product{
						Id:    product.Id,
						Stock: -p.Quantity,
					})
				}
				err = service.CustomerRepository.DeleteTransaction(ctx, res.CustomField1, res.OrderID)
				helper.PanicIfError(err)

			} else {
				continue
			}
		}
	case "failed":
		err = service.CustomerRepository.DeleteTransaction(ctx, res.CustomField1, res.OrderID)
		helper.PanicIfError(err)
	default:
		panic("not found")
	}
}
