package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"weplant-backend/model/web"
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
		if len(strings.Split(header, " ")) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   "invalid token",
			})
			return
		}
		token := strings.Split(header, " ")[1]

		fmt.Println(token)

		payload, err := middleware.JWTService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   err.Error(),
			})
			return
		}

		if role == "customer" {
			if payload.Role != "customer" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "Unauthorized",
					Data:   "You don't have permission to access",
				})
				return
			}
		} else if role == "merchant" {
			if payload.Role != "merchant" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "Unauthorized",
					Data:   "You don't have permission to access",
				})
				return
			}
		} else if role == "admin" {
			if payload.Role != "admin" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "Unauthorized",
					Data:   "You don't have permission to access",
				})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   "You don't have permission to access",
			})
			return
		}
		c.Next()

	}
}
