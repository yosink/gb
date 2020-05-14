package controllers

import (
	"blog/comm"
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

type AddArticleRequest struct {
	CID             uint   `form:"cid" json:"cid" binding:"required,numeric"`
	UserID          uint   `form:"user_id" json:"userId" binding:"required,numeric"`
	Slug            string `form:"slug" json:"slug" binding:"required,max=50,min=3"`
	Title           string `form:"title" json:"title" binding:"required,max=200,min=3"`
	Subtitle        string `form:"subtitle" json:"subtitle" binding:"max=50,omitempty"`
	Content         string `form:"content" json:"content" binding:"required"`
	PageImage       string `form:"page_image" json:"pageImage"  binding:"required,url"`
	MetaDescription string `form:"meta_description" json:"metaDescription" binding:"max=200,omitempty"`
	Recommend       uint8  `form:"recommend" json:"recommend" binding:"required"`
	Sort            int    `form:"sort" json:"sort" binding:"numeric,omitempty"`
	ViewCount       int    `form:"view_count" json:"viewCount" binding:"numeric"`
}

func AddArticle(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     AddArticleRequest
		httpCode = http.StatusOK
		errCode  = e.Success
		data     interface{}
	)
	service := services.NewArticleService()
	_, err := comm.ValidateBind(c, &form)
	if err != nil {
		appG.ResponseError(e.InvalidParams, e.InvalidParams, err.Error())
		return
	}
	// execute add
	article := map[string]interface{}{
		"cid":              form.CID,
		"user_id":          form.UserID,
		"slug":             form.Slug,
		"title":            form.Title,
		"subtitle":         form.Subtitle,
		"content":          form.Content,
		"page_image":       form.PageImage,
		"meta_description": form.MetaDescription,
		"recommend":        &form.Recommend,
		"sort":             form.Sort,
		"view_count":       form.ViewCount,
	}

	err = service.AddArticle(article)
	if err != nil {
		blogger.Error("article create error:", err)
		httpCode = http.StatusInternalServerError
		errCode = e.ErrorArticleCreate
	}
	appG.Response(httpCode, errCode, data)
}
