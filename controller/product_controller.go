package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type ProductController interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	UpdateMainImage(c *gin.Context)
	PushImageIntoImages(c *gin.Context)
	PullImageFromImages(c *gin.Context)
	Delete(c *gin.Context)
}

type productControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productControllerImpl{
		ProductService: productService,
	}
}

func (controller *productControllerImpl) Create(c *gin.Context) {
	ctx := c.Request.Context()

	merchantId := c.PostForm("merchant_id")
	name := c.PostForm("name")
	description := c.PostForm("description")
	price, err := strconv.Atoi(c.PostForm("price"))
	helper.PanicIfError(err)
	stock, err := strconv.Atoi(c.PostForm("stock"))
	helper.PanicIfError(err)
	categories := c.PostFormArray("categories")

	// main image
	image, err := c.FormFile("image")
	helper.PanicIfError(err)
	if image == nil {
		panic(errors.New("no file").Error())
	}
	src, err := image.Open()
	helper.PanicIfError(err)
	defer src.Close()
	filename := helper.GetFileName(image.Filename)

	// images
	form, err := c.MultipartForm()
	helper.PanicIfError(err)
	images := form.File["images"]

	var imagesCreateRequest []*web.ImageCreateRequest
	for _, img := range images {
		file, err := img.Open()
		helper.PanicIfError(err)
		file.Close()
		filename := helper.GetFileName(img.Filename)
		imagesCreateRequest = append(imagesCreateRequest, &web.ImageCreateRequest{
			FileName: filename,
			URL:      file,
		})
	}

	var categoriesCreateRequest []*web.ProductCategoryCreateRequest
	for _, category := range categories {
		categoriesCreateRequest = append(categoriesCreateRequest, &web.ProductCategoryCreateRequest{
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
			URL:      src,
		},
		Images:     imagesCreateRequest,
		Categories: categoriesCreateRequest,
	})

	c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   res,
	})
}

func (controller *productControllerImpl) FindById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("productId")

	res := controller.ProductService.FindById(ctx, id)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *productControllerImpl) FindAll(c *gin.Context) {
	ctx := c.Request.Context()

	res := controller.ProductService.FindAll(ctx)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *productControllerImpl) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("productId")

	var productUpdateRequest web.ProductUpdateRequest
	errBind := c.ShouldBindJSON(&productUpdateRequest)
	helper.PanicIfError(errBind)
	productUpdateRequest.Id = id
	productUpdateRequest.UpdatedAt = helper.GetTimeNow()

	res := controller.ProductService.Update(ctx, productUpdateRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *productControllerImpl) UpdateMainImage(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("productId")

	image, errorFormFile := c.FormFile("image")

	src, err := image.Open()
	helper.PanicIfError(err)
	defer src.Close()

	helper.PanicIfError(errorFormFile)
	filename := helper.GetFileName(image.Filename)

	res := controller.ProductService.UpdateMainImage(ctx, web.ProductUpdateImageRequest{
		Id:        id,
		UpdatedAt: helper.GetTimeNow(),
		MainImage: &web.ImageUpdateRequest{
			FileName: filename,
			URL:      src,
		},
	})
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *productControllerImpl) PullImageFromImages(c *gin.Context) {
	ctx := c.Request.Context()
	productId := c.Param("productId")
	imageId := c.Param("imageId")

	controller.ProductService.PullImageFromImages(ctx, productId, imageId)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (controller *productControllerImpl) PushImageIntoImages(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("productId")

	form, errorMultipartForm := c.MultipartForm()
	helper.PanicIfError(errorMultipartForm)
	images := form.File["images"]

	var imageCreateRequest []web.ImageCreateRequest
	for _, image := range images {
		src, err := image.Open()
		helper.PanicIfError(err)
		src.Close()
		filename := helper.GetFileName(image.Filename)
		imageCreateRequest = append(imageCreateRequest, web.ImageCreateRequest{
			FileName: filename,
			URL:      src,
		})
	}

	controller.ProductService.PushImageIntoImages(ctx, id, imageCreateRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (controller *productControllerImpl) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("productId")

	controller.ProductService.Delete(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
