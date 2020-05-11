package controllers

import (
	"blog/services"

	"github.com/gin-gonic/gin"
)

func GetArticles(ctx *gin.Context) {
	articleService := services.NewArticleService()
	list, err := articleService.Gets()
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "server error",
			"data":    nil,
		})
		return
	}
	ctx.JSON(200, list)
}
