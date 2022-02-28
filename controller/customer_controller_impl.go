package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	var customerCreateRequest web.CustomerCreateRequest
	helper.ReadFromRequestBody(request, &customerCreateRequest)

	customerCreateRequest.CreatedAt = helper.GetTimeNow()
	customerCreateRequest.UpdatedAt = helper.GetTimeNow()

	res := controller.CustomerService.Create(ctx, customerCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	res := controller.CustomerService.FindById(ctx, customerId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindCartById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	res := controller.CustomerService.FindCartById(ctx, customerId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindTransactionById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	res := controller.CustomerService.FindTransactionById(ctx, customerId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindOrderById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	res := controller.CustomerService.FindOrderById(ctx, customerId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	var customerUpdateRequest web.CustomerUpdateRequest
	helper.ReadFromRequestBody(request, &customerUpdateRequest)
	customerUpdateRequest.Id = customerId
	customerUpdateRequest.UpdatedAt = helper.GetTimeNow()

	res := controller.CustomerService.Update(ctx, customerUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) UpdateMainImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	file, fileHeader, err := request.FormFile("image")
	helper.PanicIfError(err)

	filename := helper.GetFileName(fileHeader.Filename)

	res := controller.CustomerService.UpdateMainImage(ctx, web.CustomerUpdateImageRequest{
		Id:        customerId,
		UpdatedAt: helper.GetTimeNow(),
		MainImage: &web.ImageUpdateRequest{
			FileName: filename,
			URL:      file,
		},
	})
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	customerId := params.ByName("customerId")

	controller.CustomerService.Delete(ctx, customerId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
