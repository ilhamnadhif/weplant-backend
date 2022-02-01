package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go/coreapi"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type TransactionController interface {
	Create(c *gin.Context)
	Cancel(c *gin.Context)
	Callback(c *gin.Context)
}

type transactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &transactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *transactionControllerImpl) Create(c *gin.Context) {
	ctx := c.Request.Context()
	customerId := c.Param("customerId")

	var addressCreateRequest web.AddressCreateRequest
	err := c.ShouldBindJSON(&addressCreateRequest)
	helper.PanicIfError(err)

	res := controller.TransactionService.Create(ctx, web.TransactionCreateRequest{
		CreatedAt:  helper.GetTimeNow(),
		UpdatedAt:  helper.GetTimeNow(),
		CustomerId: customerId,
		Address:    &addressCreateRequest,
	})

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *transactionControllerImpl) Cancel(c *gin.Context) {
	ctx := c.Request.Context()
	customerId := c.Param("customerId")
	transactionId := c.Param("transactionId")

	controller.TransactionService.Cancel(ctx, customerId, transactionId)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (controller *transactionControllerImpl) Callback(c *gin.Context) {
	ctx := c.Request.Context()
	var cb coreapi.TransactionStatusResponse

	err := c.ShouldBindJSON(&cb)
	helper.PanicIfError(err)

	controller.TransactionService.Callback(ctx, cb)
	c.AbortWithStatus(http.StatusOK)
}
