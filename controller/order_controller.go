package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go/coreapi"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type OrderController interface {
	CheckoutFromCart(c *gin.Context)
	CallbackTransaction(c *gin.Context)
}

type orderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *orderControllerImpl) CheckoutFromCart(c *gin.Context) {
	ctx := c.Request.Context()
	customerId := c.Param("customerId")

	var addressCreateRequest web.AddressCreateRequest
	err := c.ShouldBindJSON(&addressCreateRequest)
	helper.PanicIfError(err)

	res := controller.OrderService.CheckoutFromCart(ctx, web.OrderCreateRequest{
		CustomerId: customerId,
		Address:    &addressCreateRequest,
	})
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *orderControllerImpl) CallbackTransaction(c *gin.Context) {
	ctx := c.Request.Context()

	var cb coreapi.TransactionStatusResponse

	err := c.ShouldBindJSON(&cb)
	helper.PanicIfError(err)

	res := controller.OrderService.CallbackTransaction(ctx, cb)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}
