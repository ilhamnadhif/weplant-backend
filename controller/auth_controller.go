package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthController interface {
	LoginCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	LoginMerchant(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
