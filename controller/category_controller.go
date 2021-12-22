package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type CategoryController interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	UpdateMainImage(c *gin.Context)
	Delete(c *gin.Context)
}

type categoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *categoryControllerImpl) Create(c *gin.Context) {
	ctx := c.Request.Context()

	name := c.PostForm("name")

	image, errorFormFile := c.FormFile("image")

	src, err := image.Open()
	helper.PanicIfError(err)
	defer src.Close()

	helper.PanicIfError(errorFormFile)
	filename := helper.GetFileName(image.Filename)

	res := controller.CategoryService.Create(ctx, web.CategoryCreateRequest{
		Name: name,
		MainImage: &web.ImageCreateRequest{
			FileName: filename,
		},
	}, src)
	c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   res,
	})
}

func (controller *categoryControllerImpl) FindById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("categoryId")

	res := controller.CategoryService.FindById(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *categoryControllerImpl) FindAll(c *gin.Context) {
	ctx := c.Request.Context()

	res := controller.CategoryService.FindAll(ctx)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *categoryControllerImpl) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("categoryId")

	var categoryUpdateRequest web.CategoryUpdateRequest
	errBind := c.ShouldBindJSON(&categoryUpdateRequest)
	if errBind != nil {
		panic(helper.IfValidationError(errBind))
	}
	categoryUpdateRequest.Id = id

	res := controller.CategoryService.Update(ctx, categoryUpdateRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *categoryControllerImpl) UpdateMainImage(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("categoryId")

	image, errorFormFile := c.FormFile("image")

	src, err := image.Open()
	helper.PanicIfError(err)
	defer src.Close()

	helper.PanicIfError(errorFormFile)
	filename := helper.GetFileName(image.Filename)

	res := controller.CategoryService.UpdateMainImage(ctx, web.CategoryUpdateImageRequest{
		Id:        id,
		MainImage: &web.ImageUpdateRequest{
			FileName: filename,
		},
	}, src)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *categoryControllerImpl) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("categoryId")

	controller.CategoryService.Delete(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
