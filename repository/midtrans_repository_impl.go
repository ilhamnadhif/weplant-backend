package repository

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"os"
	"weplant-backend/helper"
)

type MidtransRepositoryImpl struct {
	ServerKey string
}

func NewMidtransRepository(serverKey string) MidtransRepository {
	return &MidtransRepositoryImpl{
		ServerKey: serverKey,
	}
}

func (repository *MidtransRepositoryImpl) CreateTransaction(req coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
	var c coreapi.Client
	c.New(repository.ServerKey, helper.MidtransEnvType(os.Getenv("GO_ENV")))

	return c.ChargeTransaction(&req)
}

func (repository *MidtransRepositoryImpl) CancelTransaction(orderId string) (*coreapi.CancelResponse, *midtrans.Error) {
	var c coreapi.Client
	c.New(repository.ServerKey, helper.MidtransEnvType(os.Getenv("GO_ENV")))

	return c.CancelTransaction(orderId)
}

func (repository *MidtransRepositoryImpl) CheckTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {

	var c coreapi.Client
	c.New(repository.ServerKey, helper.MidtransEnvType(os.Getenv("GO_ENV")))

	return c.CheckTransaction(orderId)
}
