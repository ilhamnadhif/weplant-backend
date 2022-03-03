package repository_mock

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/stretchr/testify/mock"
)

type MidtransRepositoryMock struct {
	Mock mock.Mock
}

func (repository *MidtransRepositoryMock) CreateTransaction(req coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
	arguments := repository.Mock.Called(req)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(*midtrans.Error)
	}

	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(*midtrans.Error)
	} else {
		return arguments.Get(0).(*coreapi.ChargeResponse), nil
	}
}

func (repository *MidtransRepositoryMock) CancelTransaction(orderId string) (*coreapi.CancelResponse, *midtrans.Error) {

	arguments := repository.Mock.Called(orderId)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(*midtrans.Error)
	}

	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(*midtrans.Error)
	} else {
		return arguments.Get(0).(*coreapi.CancelResponse), nil
	}
}

func (repository *MidtransRepositoryMock) CheckTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {
	arguments := repository.Mock.Called(orderId)

	if arguments.Get(1) != nil {
		return nil, arguments.Get(1).(*midtrans.Error)
	}

	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(*midtrans.Error)
	} else {
		return arguments.Get(0).(*coreapi.TransactionStatusResponse), nil
	}
}
