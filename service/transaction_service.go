package service

import (
	"context"
	"github.com/midtrans/midtrans-go/coreapi"
	"weplant-backend/model/web"
)

type TransactionService interface {
	Create(ctx context.Context, request web.TransactionCreateRequest) web.TransactionCreateRequestResponse
	Cancel(ctx context.Context, customerId string, transactionId string)
	Callback(ctx context.Context, request coreapi.TransactionStatusResponse)
}
