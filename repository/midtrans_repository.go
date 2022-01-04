package repository

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"os"
	"weplant-backend/config"
	"weplant-backend/helper"
)

type MidtransRepository interface {
	Checkout(request snap.Request) (*snap.Response, *midtrans.Error)
	CheckTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
}

type midtransRepositoryImpl struct {
}

func NewMidtransRepository() MidtransRepository {
	return &midtransRepositoryImpl{}
}

func (repository *midtransRepositoryImpl) Checkout(request snap.Request) (*snap.Response, *midtrans.Error) {
	key := config.GetMidtransKey()

	var s snap.Client
	s.New(key, helper.MidtransEnvType(os.Getenv("MIDTRANS_ENV_TYPE")))

	res, err := s.CreateTransaction(&request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repository *midtransRepositoryImpl) CheckTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {
	key := config.GetMidtransKey()

	var c coreapi.Client
	c.New(key, helper.MidtransEnvType(os.Getenv("MIDTRANS_ENV_TYPE")))

	return c.CheckTransaction(orderId)
}

