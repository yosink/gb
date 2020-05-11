package routers

import "github.com/gin-gonic/gin"

func LoadWebRoutes(r *gin.Engine) {
	wap := r.Group("/wap")
	{
		wap.GET("/index", func(ctx *gin.Context) {
			ctx.String(200, "wap index")
		})
	}
}
