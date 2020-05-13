package controllers

import (
	"blog/comm"
	blogger "blog/logging"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/services"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"

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
	CID             int    `form:"cid" json:"cid" valid:",int~cid类型错误,required~缺少cid" binding:"required"`
	UserID          uint   `form:"user_id" json:"userId" valid:"int~user_id类型错误,required~缺少user_id" binding:"required"`
	Slug            string `form:"slug" json:"slug" valid:"required,stringlength(3|20),slug_unique" binding:"required"`
	Title           string `form:"title" json:"title" valid:"required,stringlength(2|100)" binding:"required"`
	Subtitle        string `form:"subtitle" json:"subtitle" valid:",stringlength(2|50),optional"`
	Content         string `form:"content" json:"contIent" valid:"required" binding:"required"`
	PageImage       string `form:"page_image" json:"pageImage" valid:"required,url" binding:"required"`
	MetaDescription string `form:"meta_description" json:"metaDescription" valid:"optional"`
	Recommend       uint8  `form:"recommend" json:"recommend" valid:"range(0|1),required" binding:"required"`
	Sort            int    `form:"sort" json:"sort "valid:"range(0|255),optional"`
	ViewCount       int    `form:"view_count" json:"viewCount" valid:"int"`
}

func AddArticle(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     AddArticleRequest
		httpCode = e.Success
		errCode  int
		data     interface{}
	)
	service := services.NewArticleService()
	govalidator.TagMap["slug_unique"] = func(str string) bool {
		exists, _ := service.Exists(map[string]interface{}{"slug": str})
		return !exists
	}
	_, err := comm.BindAndValid(c, &form)
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
		errCode = e.ErrorArticleCreate
	}
	appG.Response(httpCode, errCode, data)
}
