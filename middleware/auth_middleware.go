package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"weplant-backend/exception"
	"weplant-backend/pkg"
)


func  AuthMiddleware(handle httprouter.Handle, role string) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		header := request.Header.Get("Authorization")
		if len(strings.Split(header, " ")) != 2 {
			panic(exception.NewUnauthorizedError("auth header is invalid"))
			return
		}
		token := strings.Split(header, " ")[1]

		payload, err := pkg.ValidateToken(token)
		if err != nil {
			panic(exception.NewUnauthorizedError(err.Error()))
			return
		}

		switch role {
		case "admin":
			if payload.Role != "admin" {
				panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
				return
			} else {
				handle(writer, request, params)
			}
		case "merchant":
			if payload.Role != "merchant" {
				panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
				return
			} else {
				handle(writer, request, params)
			}
		case "customer":
			if payload.Role != "customer" {
				panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
				return
			} else {
				handle(writer, request, params)
			}
		default:
			panic(exception.NewUnauthorizedError("you don't have permission to access this resource"))
			return
		}
	}
}
