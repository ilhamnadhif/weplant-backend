package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type CartController interface {
	PushProductToCart(c *gin.Context)
	UpdateProductQuantity(c *gin.Context)
	PullProductFromCart(c *gin.Context)
}

type cartControllerImpl struct {
	CartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &cartControllerImpl{
		CartService: cartService,
	}
}

func (controller *cartControllerImpl) PushProductToCart(c *gin.Context) {
	ctx := c.Request.Context()
	customerId := c.Param("customerId")

	var cartRequest web.CartProductCreateRequest
	errBind := c.ShouldBindJSON(&cartRequest)
	helper.PanicIfError(errBind)
	cartRequest.CreatedAt = helper.GetTimeNow()
	cartRequest.UpdatedAt = helper.GetTimeNow()
	cartRequest.CustomerId = customerId
	cartRequest.Quantity = 1

	res := controller.CartService.PushProductToCart(ctx, cartRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *cartControllerImpl) UpdateProductQuantity(c *gin.Context) {
	ctx := c.Request.Context()
	customerId := c.Param("customerId")
	productId := c.Param("productId")

	var cartRequest web.CartProductUpdateRequest
	errBind := c.ShouldBindJSON(&cartRequest)
	helper.PanicIfError(errBind)
	cartRequest.UpdatedAt = helper.GetTimeNow()
	cartRequest.CustomerId = customerId
	cartRequest.ProductId = productId

	res := controller.CartService.UpdateProductQuantity(ctx, cartRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *cartControllerImpl) PullProductFromCart(c *gin.Context) {
	ctx := c.Request.Context()
	customerId := c.Param("customerId")
	productId := c.Param("productId")

	controller.CartService.PullProductFromCart(ctx, customerId, productId)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
