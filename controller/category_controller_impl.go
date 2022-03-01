package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	var categoryCreateRequest web.CategoryCreateRequest
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	res := controller.CategoryService.Create(ctx, web.CategoryCreateRequest{
		CreatedAt: helper.GetTimeNow(),
		UpdatedAt: helper.GetTimeNow(),
		Name:      categoryCreateRequest.Name,
		Slug:      helper.SlugGenerate(categoryCreateRequest.Name),
	})
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	categoryId := params.ByName("categoryId")

	res := controller.CategoryService.FindById(ctx, categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	res := controller.CategoryService.FindAll(ctx)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
