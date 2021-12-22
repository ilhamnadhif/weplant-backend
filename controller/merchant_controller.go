package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type MerchantController interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	UpdateMainImage(c *gin.Context)
	Delete(c *gin.Context)
}

type merchantControllerImpl struct {
	MerchantService service.MerchantService
}

func NewMerchantController(merchantService service.MerchantService) MerchantController {
	return &merchantControllerImpl{
		MerchantService: merchantService,
	}
}

func (controller *merchantControllerImpl) Create(c *gin.Context) {
	ctx := c.Request.Context()

	email := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	city := c.PostForm("city")
	province := c.PostForm("province")
	country := c.PostForm("country")
	postalCode, errPosCod := strconv.Atoi(c.PostForm("postal_code"))
	helper.PanicIfError(errPosCod)
	latitude, errlat := strconv.ParseFloat(c.PostForm("latitude"), 64)
	helper.PanicIfError(errlat)
	longitude, errLong := strconv.ParseFloat(c.PostForm("longitude"), 64)
	helper.PanicIfError(errLong)

	image, errorFormFile := c.FormFile("image")

	src, err := image.Open()
	helper.PanicIfError(err)
	defer src.Close()

	helper.PanicIfError(errorFormFile)
	filename := helper.GetFileName(image.Filename)

	res := controller.MerchantService.Create(ctx, web.MerchantCreateRequest{
		Email:    email,
		Password: password,
		Name:     name,
		Phone:    phone,
		MainImage: &web.ImageCreateRequest{
			FileName: filename,
		},
		Address: &web.AddressCreateRequest{
			Address:    address,
			City:       city,
			Province:   province,
			Country:    country,
			PostalCode: postalCode,
			Latitude:   latitude,
			Longitude:  longitude,
		},
	}, src)
	c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   res,
	})
}

func (controller *merchantControllerImpl) FindById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("merchantId")

	res := controller.MerchantService.FindById(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *merchantControllerImpl) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("merchantId")

	var merchantUpdateRequest web.MerchantUpdateRequest
	merchantUpdateRequest.Id = id

	errBind := c.ShouldBindJSON(&merchantUpdateRequest)
	if errBind != nil {
		panic(helper.IfValidationError(errBind))
	}
	res := controller.MerchantService.Update(ctx, merchantUpdateRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *merchantControllerImpl) UpdateMainImage(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("merchantId")

	image, errorFormFile := c.FormFile("image")

	src, err := image.Open()
	helper.PanicIfError(err)
	defer src.Close()

	helper.PanicIfError(errorFormFile)
	filename := helper.GetFileName(image.Filename)

	res := controller.MerchantService.UpdateMainImage(ctx, web.MerchantUpdateImageRequest{
		Id: id,
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

func (controller *merchantControllerImpl) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("merchantId")

	controller.MerchantService.Delete(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
