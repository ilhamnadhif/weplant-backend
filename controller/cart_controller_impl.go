package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type CartControllerImpl struct {
	CartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &CartControllerImpl{
		CartService: cartService,
	}
}

func (controller *CartControllerImpl) PushProductToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	var cartRequest web.CartProductCreateRequest
	helper.ReadFromRequestBody(request, &cartRequest)
	cartRequest.CustomerId = customerId
	cartRequest.Quantity = 1

	res := controller.CartService.PushProductToCart(ctx, cartRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) UpdateProductQuantity(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")
	productId := params.ByName("productId")

	var cartRequest web.CartProductUpdateRequest
	helper.ReadFromRequestBody(request, &cartRequest)
	cartRequest.UpdatedAt = helper.GetTimeNow()
	cartRequest.CustomerId = customerId
	cartRequest.ProductId = productId

	res := controller.CartService.UpdateProductQuantity(ctx, cartRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) PullProductFromCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")
	productId := params.ByName("productId")

	controller.CartService.PullProductFromCart(ctx, customerId, productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
