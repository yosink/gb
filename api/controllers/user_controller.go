package controllers

import (
	"blog/pkg/app"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	g := app.Gin{C: c}
	g.Response(200, 200, "this is me")
}
