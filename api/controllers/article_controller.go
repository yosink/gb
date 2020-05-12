package controllers

import (
	blogger "blog/logging"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/services"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

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

func GetArticle(c *gin.Context) {
	var code int
	var data interface{}
	code = e.Success

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		code = e.InvalidParams
	} else {
		articleService := services.NewArticleService()
		article, err := articleService.GetByID(id)
		if err != nil {
			blogger.Error(err)
			if gorm.IsRecordNotFoundError(err) {
				code = e.ArticleNotFound
			} else {
				code = e.ERROR
			}
		} else {
			data = article
		}
	}
	a := app.Gin{C: c}
	a.Response(http.StatusOK, code, data)
}
