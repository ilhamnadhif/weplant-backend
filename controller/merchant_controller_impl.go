package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type MerchantControllerImpl struct {
	MerchantService service.MerchantService
}

func NewMerchantController(merchantService service.MerchantService) MerchantController {
	return &MerchantControllerImpl{
		MerchantService: merchantService,
	}
}

func (controller *MerchantControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	email := request.PostFormValue("email")
	password := request.PostFormValue("password")
	name := request.PostFormValue("name")
	phone := request.PostFormValue("phone")
	address := request.PostFormValue("address")
	city := request.PostFormValue("city")
	province := request.PostFormValue("province")
	country := request.PostFormValue("country")
	postalCode := request.PostFormValue("postal_code")

	file, fileHeader, err := request.FormFile("image")

	helper.PanicIfError(err)

	filename := helper.GetFileName(fileHeader.Filename)

	res := controller.MerchantService.Create(ctx, web.MerchantCreateRequest{
		CreatedAt: helper.GetTimeNow(),
		UpdatedAt: helper.GetTimeNow(),
		Email:     email,
		Password:  password,
		Name:      name,
		Slug:      helper.SlugGenerate(name),
		Phone:     phone,
		MainImage: &web.ImageCreateRequest{
			FileName: filename,
			URL:      file,
		},
		Address: &web.AddressCreateRequest{
			Address:    address,
			City:       city,
			Province:   province,
			Country:    country,
			PostalCode: postalCode,
		},
	})
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	merchantId := params.ByName("merchantId")

	res := controller.MerchantService.FindById(ctx, merchantId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) FindManageOrderById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	merchantId := params.ByName("merchantId")

	res := controller.MerchantService.FindManageOrderById(ctx, merchantId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	merchantId := params.ByName("merchantId")

	var merchantUpdateRequest web.MerchantUpdateRequest

	helper.ReadFromRequestBody(request, &merchantUpdateRequest)
	merchantUpdateRequest.Id = merchantId
	merchantUpdateRequest.UpdatedAt = helper.GetTimeNow()

	res := controller.MerchantService.Update(ctx, merchantUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) UpdateMainImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	merchantId := params.ByName("merchantId")

	file, fileHeader, err := request.FormFile("image")

	helper.PanicIfError(err)
	filename := helper.GetFileName(fileHeader.Filename)

	res := controller.MerchantService.UpdateMainImage(ctx, web.MerchantUpdateImageRequest{
		Id:        merchantId,
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

func (controller *MerchantControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	merchantId := params.ByName("merchantId")

	controller.MerchantService.Delete(ctx, merchantId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
