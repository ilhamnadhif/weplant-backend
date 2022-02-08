package repository

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransRepository interface {
	CreateTransaction(req coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error)
	CancelTransaction(orderId string) (*coreapi.CancelResponse, *midtrans.Error)
	CheckTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
}
