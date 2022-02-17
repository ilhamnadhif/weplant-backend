package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) LoginCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	var loginRequest web.LoginRequest
	helper.ReadFromRequestBody(request, &loginRequest)

	customer := controller.AuthService.LoginCustomer(ctx, loginRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customer,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) LoginMerchant(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	var loginRequest web.LoginRequest
	helper.ReadFromRequestBody(request, &loginRequest)

	merchant := controller.AuthService.LoginMerchant(ctx, loginRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   merchant,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) LoginAdmin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	var loginRequest web.LoginRequest
	helper.ReadFromRequestBody(request, &loginRequest)

	admin := controller.AuthService.LoginAdmin(ctx, loginRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   admin,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
