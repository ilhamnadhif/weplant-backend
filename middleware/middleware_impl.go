package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"weplant-backend/exception"
	"weplant-backend/service"
)

type MiddlewareImpl struct {
	JWTService service.JWTService
}

func NewMiddleware(jwtService service.JWTService) Middleware {
	return &MiddlewareImpl{
		JWTService: jwtService,
	}
}

func (middleware *MiddlewareImpl) AuthMiddleware(handle httprouter.Handle, role string) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		header := request.Header.Get("Authorization")
		if len(strings.Split(header, " ")) != 2 {
			panic(exception.NewUnauthorizedError("auth header is invalid"))
		}
		token := strings.Split(header, " ")[1]

		payload, err := middleware.JWTService.ValidateToken(token)
		if err != nil {
			panic(exception.NewUnauthorizedError(err.Error()))
		}

		switch role {
		case "admin":
			if payload.Role != "admin" {
				panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
			} else {
				handle(writer, request, params)
			}
		case "merchant":
			if payload.Role != "merchant" {
				panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
			} else {
				handle(writer, request, params)
			}
		case "customer":
			if payload.Role != "customer" {
				panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
			} else {
				handle(writer, request, params)
			}
		default:
			panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
		}
	}
}
