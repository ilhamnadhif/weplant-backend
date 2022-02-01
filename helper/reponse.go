package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weplant-backend/model/web"
)

func ResponseUnauthorized(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
		Data:   message,
	})
}
