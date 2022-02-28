package controller

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	err := request.ParseMultipartForm(32 << 20)
	helper.PanicIfError(err)

	merchantId := request.PostFormValue("merchant_id")
	name := request.PostFormValue("name")
	description := request.PostFormValue("description")
	price, err := strconv.Atoi(request.PostFormValue("price"))
	helper.PanicIfError(err)
	stock, err := strconv.Atoi(request.PostFormValue("stock"))
	helper.PanicIfError(err)

	// main image
	file, fileHeader, err := request.FormFile("image")
	helper.PanicIfError(err)
	if file == nil {
		panic(errors.New("no file").Error())
	}
	filename := helper.GetFileName(fileHeader.Filename)

	// images
	images := request.MultipartForm.File["images"]
	var imagesCreateRequest []web.ImageCreateRequest
	for _, img := range images {
		file, err := img.Open()
		helper.PanicIfError(err)
		file.Close()
		imagesCreateRequest = append(imagesCreateRequest, web.ImageCreateRequest{
			FileName: helper.GetFileName(img.Filename),
			URL:      file,
		})
	}

	// categories
	categories := request.PostForm["categories"]
	var categoriesCreateRequest []web.ProductCategoryCreateRequest
	for _, category := range categories {
		categoriesCreateRequest = append(categoriesCreateRequest, web.ProductCategoryCreateRequest{
			CategoryId: category,
		})
	}

	res := controller.ProductService.Create(ctx, web.ProductCreateRequest{
		CreatedAt:   helper.GetTimeNow(),
		UpdatedAt:   helper.GetTimeNow(),
		MerchantId:  merchantId,
		Name:        name,
		Slug:        helper.SlugGenerate(name),
		Description: description,
		Price:       price,
		Stock:       stock,
		MainImage: &web.ImageCreateRequest{
			FileName: filename,
			URL:      file,
		},
		Images:     imagesCreateRequest,
		Categories: categoriesCreateRequest,
	})

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	productId := params.ByName("productId")

	res := controller.ProductService.FindById(ctx, productId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	var page int
	var perPage int

	queryPage := request.URL.Query().Get("page")
	queryPerPage := request.URL.Query().Get("perPage")

	if queryPage == "" {
		page = 1
	} else {
		a, err := strconv.Atoi(queryPage)
		helper.PanicIfError(err)
		page = a
	}

	if queryPerPage == "" {
		perPage = 10
	} else {
		a, err := strconv.Atoi(queryPerPage)
		helper.PanicIfError(err)
		perPage = a
	}

	search := request.URL.Query().Get("search")

	if search != "" {
		res := controller.ProductService.FindAllWithSearch(ctx, search, page, perPage)
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   res,
		}
		helper.WriteToResponseBody(writer, webResponse)
	} else {
		res := controller.ProductService.FindAll(ctx, page, perPage)
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   res,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}

}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	productId := params.ByName("productId")

	var productUpdateRequest web.ProductUpdateRequest

	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productUpdateRequest.Id = productId
	productUpdateRequest.UpdatedAt = helper.GetTimeNow()

	res := controller.ProductService.Update(ctx, productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) UpdateMainImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	productId := params.ByName("productId")

	file, fileHeader, err := request.FormFile("image")
	helper.PanicIfError(err)

	filename := helper.GetFileName(fileHeader.Filename)

	res := controller.ProductService.UpdateMainImage(ctx, web.ProductUpdateImageRequest{
		Id:        productId,
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

func (controller *ProductControllerImpl) PullImageFromImages(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	productId := params.ByName("productId")
	imageId := params.ByName("imageId")

	controller.ProductService.PullImageFromImages(ctx, productId, imageId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) PushImageIntoImages(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	productId := params.ByName("productId")

	err := request.ParseMultipartForm(32 << 20)
	helper.PanicIfError(err)

	images := request.MultipartForm.File["images"]

	var imageCreateRequest []web.ImageCreateRequest
	for _, image := range images {
		file, err := image.Open()
		helper.PanicIfError(err)
		imageCreateRequest = append(imageCreateRequest, web.ImageCreateRequest{
			FileName: helper.GetFileName(image.Filename),
			URL:      file,
		})
	}

	controller.ProductService.PushImageIntoImages(ctx, productId, imageCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	productId := params.ByName("productId")

	controller.ProductService.Delete(ctx, productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
