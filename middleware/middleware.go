package middleware

import (
	"github.com/julienschmidt/httprouter"
)

type Middleware interface {
	AuthMiddleware(handle httprouter.Handle, role string) httprouter.Handle
}
