package routers

import (
	"blog/pkg/file"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	f, _ := file.MustOpen("gin"+time.Now().Format("2006-01-02"), "runtime")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/assets", "./public")

	LoadApiRoutes(r)
	LoadWebRoutes(r)

	return r
}
