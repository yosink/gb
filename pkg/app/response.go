package app

import (
	"blog/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errorCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errorCode,
		Message: e.GetMessage(errorCode),
		Data:    data,
	})
}

func (g *Gin) ResponseError(httpErrorCode, errorCode int, msg string) {
	g.C.JSON(httpErrorCode, Response{
		Code:    errorCode,
		Message: msg,
	})
}
