package app

import (
	"blog/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func (g *Gin) Response(httpCode, errorCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errorCode,
		Message: e.GetMessage(errorCode),
		Data:    data,
	})
}
