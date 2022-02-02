package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weplant-backend/helper"
	"weplant-backend/model/web"
	"weplant-backend/service"
)

type AuthController interface {
	LoginCustomer(c *gin.Context)
	LoginMerchant(c *gin.Context)
	LoginAdmin(c *gin.Context)
}

type authControllerImpl struct {
	AuthService service.AuthService
	JWTService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authControllerImpl{
		AuthService: authService,
		JWTService:  jwtService,
	}
}

func (controller *authControllerImpl) LoginCustomer(c *gin.Context) {
	ctx := c.Request.Context()

	var loginRequest web.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	helper.PanicIfError(err)

	customer := controller.AuthService.LoginCustomer(ctx, loginRequest)

	token := controller.JWTService.GenerateToken(web.JWTPayload{
		Id:   customer.Id,
		Role: customer.Role,
	})

	customer.Token = token

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customer,
	})
}

func (controller *authControllerImpl) LoginMerchant(c *gin.Context) {
	ctx := c.Request.Context()

	var loginRequest web.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	helper.PanicIfError(err)

	merchant := controller.AuthService.LoginMerchant(ctx, loginRequest)

	token := controller.JWTService.GenerateToken(web.JWTPayload{
		Id:   merchant.Id,
		Role: merchant.Role,
	})

	merchant.Token = token

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   merchant,
	})
}

func (controller *authControllerImpl) LoginAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	var loginRequest web.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	helper.PanicIfError(err)

	admin := controller.AuthService.LoginAdmin(ctx, loginRequest)

	token := controller.JWTService.GenerateToken(web.JWTPayload{
		Id:   admin.Id,
		Role: admin.Role,
	})

	admin.Token = token

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   admin,
	})
}
