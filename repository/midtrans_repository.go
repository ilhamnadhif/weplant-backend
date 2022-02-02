package repository

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"os"
	"weplant-backend/helper"
)

type MidtransRepository interface {
	CreateTransaction(req coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error)
	CancelTransaction(orderId string) (*coreapi.CancelResponse, *midtrans.Error)
	CheckTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
}

type midtransRepositoryImpl struct {
	ServerKey string
}

func NewMidtransRepository(serverKey string) MidtransRepository {
	return &midtransRepositoryImpl{
		ServerKey: serverKey,
	}
}

func (repository *midtransRepositoryImpl) CreateTransaction(req coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
	var c coreapi.Client
	c.New(repository.ServerKey, helper.MidtransEnvType(os.Getenv("ENV_MODE")))

	return c.ChargeTransaction(&req)
}

func (repository *midtransRepositoryImpl) CancelTransaction(orderId string) (*coreapi.CancelResponse, *midtrans.Error) {
	var c coreapi.Client
	c.New(repository.ServerKey, helper.MidtransEnvType(os.Getenv("ENV_MODE")))

	return c.CancelTransaction(orderId)
}

func (repository *midtransRepositoryImpl) CheckTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {

	var c coreapi.Client
	c.New(repository.ServerKey, helper.MidtransEnvType(os.Getenv("ENV_MODE")))

	return c.CheckTransaction(orderId)
}
