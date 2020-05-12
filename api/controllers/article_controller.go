package controllers

import (
	blogger "blog/logging"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/services"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"

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
	CID             uint   `valid:"int,required"`
	UserID          uint   `valid:"int,required"`
	Slug            string `valid:"stringlength(3|20),optional"`
	Title           string `valid:"required,stringlength(2|100)"`
	Subtitle        string `valid:"stringlength(2|50),optional"`
	Content         string `valid:"required"`
	PageImage       string `valid:"required,url"`
	MetaDescription string `valid:"optional"`
	Recommend       uint8  `valid:"range(0|1),required"`
	Sort            uint   `valid:"range(0|255),optional"`
	ViewCount       uint   `valid:"int"`
}

func AddArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddArticleRequest
		code = e.Success
	)
	_ = c.Bind(&form)
	valid.SetFieldsRequiredByDefault(true)
	_, err := valid.ValidateStruct(form)
	if err != nil {
		code = e.InvalidParams
	}
	appG.Response(http.StatusOK, code, nil)
}
