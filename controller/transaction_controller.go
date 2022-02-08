package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TransactionController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Cancel(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Callback(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
