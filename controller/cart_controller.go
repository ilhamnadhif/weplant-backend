package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CartController interface {
	PushProductToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateProductQuantity(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	PullProductFromCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
