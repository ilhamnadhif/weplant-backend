package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"weplant-backend/helper"
	"weplant-backend/service"
)

type AuthMiddleware interface {
	AuthJWT(role string) gin.HandlerFunc
}

type authMiddleWareImpl struct {
	JWTService service.JWTService
}

func NewAuthMiddleware(jwtService service.JWTService) AuthMiddleware {
	return &authMiddleWareImpl{
		JWTService: jwtService,
	}
}


func (middleware *authMiddleWareImpl) AuthJWT(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if len(strings.Split(header, " ")) != 2 {
			helper.ResponseUnauthorized(c, "auth header is invalid")
			return
		}
		token := strings.Split(header, " ")[1]

		payload, err := middleware.JWTService.ValidateToken(token)
		if err != nil {
			helper.ResponseUnauthorized(c, err.Error())
			return
		}

		switch role {
		case "admin":
			if payload.Role != "admin" {
				helper.ResponseUnauthorized(c, "you don't have permission to access this resource")
				return
			}
		case "merchant":
			if payload.Role != "merchant" {
				helper.ResponseUnauthorized(c, "you don't have permission to access this resource")
				return
			}
		case "customer":
			if payload.Role != "customer" {
				helper.ResponseUnauthorized(c, "you don't have permission to access this resource")
				return
			}
		default:
			helper.ResponseUnauthorized(c, "you don't have permission to access this resource")
			return
		}
		c.Next()
	}
}
