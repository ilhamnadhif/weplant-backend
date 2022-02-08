package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go/coreapi"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	var addressCreateRequest web.AddressCreateRequest
	helper.ReadFromRequestBody(request, &addressCreateRequest)

	res := controller.TransactionService.Create(ctx, web.TransactionCreateRequest{
		CreatedAt:  helper.GetTimeNow(),
		UpdatedAt:  helper.GetTimeNow(),
		CustomerId: customerId,
		Address:    &addressCreateRequest,
	})

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Cancel(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")
	transactionId := params.ByName("transactionId")

	controller.TransactionService.Cancel(ctx, customerId, transactionId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Callback(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	var cb coreapi.TransactionStatusResponse
	helper.ReadFromRequestBody(request, &cb)

	controller.TransactionService.Callback(ctx, cb)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
