package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CustomerController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindCartById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindTransactionById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindOrderById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateMainImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
