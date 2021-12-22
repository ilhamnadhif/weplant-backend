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
	Delete(c *gin.Context)
	PushImageIntoImages(c *gin.Context)
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
	price, errPrice := strconv.Atoi(c.PostForm("price"))
	helper.PanicIfError(errPrice)
	quantity, errQuantity := strconv.Atoi(c.PostForm("quantity"))
	helper.PanicIfError(errQuantity)
	categories := c.PostFormArray("categories")

	image, errorFormFile := c.FormFile("image")

	if image == nil {
		panic(errors.New("no file").Error())
	}
	src, err := image.Open()
	helper.PanicIfError(err)
	defer src.Close()

	helper.PanicIfError(errorFormFile)
	filename := helper.GetFileName(image.Filename)

	var categoriesCreateRequest []*web.ProductCategoryCreateRequest
	for _, category := range categories {
		categoriesCreateRequest = append(categoriesCreateRequest, &web.ProductCategoryCreateRequest{
			CategoryId: category,
		})
	}

	res := controller.ProductService.Create(ctx, web.ProductCreateRequest{
		MerchantId:  merchantId,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		MainImage: &web.ImageCreateRequest{
			FileName: filename,
		},
		Categories: categoriesCreateRequest,
	}, src)

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
	panic("implement me")
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
func (controller *productControllerImpl) PushImageIntoImages(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("productId")

	form, errorMultipartForm := c.MultipartForm()
	helper.PanicIfError(errorMultipartForm)
	images := form.File["images"]

	for _, image := range images {
		src, err := image.Open()
		helper.PanicIfError(err)
		defer src.Close()
		filename := helper.GetFileName(image.Filename)
		controller.ProductService.PushImageIntoImages(ctx, id, web.ImageCreateRequest{
			FileName: filename,
		}, src)
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
