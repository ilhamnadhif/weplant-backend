package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type CustomerController interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
	FindCartById(c * gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type customerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController  {
	return &customerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *customerControllerImpl) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var customerCreateRequest web.CustomerCreateRequest
	errBind := c.ShouldBindJSON(&customerCreateRequest)
	helper.PanicIfError(errBind)

	customerCreateRequest.CreatedAt = helper.GetTimeNow()
	customerCreateRequest.UpdatedAt = helper.GetTimeNow()

	res := controller.CustomerService.Create(ctx, customerCreateRequest)
	c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   res,
	})
}

func (controller *customerControllerImpl) FindById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("customerId")

	res := controller.CustomerService.FindById(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *customerControllerImpl) FindCartById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("customerId")

	res := controller.CustomerService.FindCartById(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *customerControllerImpl) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("customerId")

	var customerUpdateRequest web.CustomerUpdateRequest
	errBind := c.ShouldBindJSON(&customerUpdateRequest)
	helper.PanicIfError(errBind)
	customerUpdateRequest.Id = id
	customerUpdateRequest.UpdatedAt = helper.GetTimeNow()

	res := controller.CustomerService.Update(ctx, customerUpdateRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *customerControllerImpl) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("customerId")

	controller.CustomerService.Delete(ctx, id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
