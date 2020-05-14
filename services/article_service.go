package services

import (
	blogger "blog/logging"
	"blog/models"
	"blog/pkg/gredis"
	"blog/repositories"
	"encoding/json"
	"strconv"
)

type IArticleService interface {
	Gets() ([]*models.ArticleCopy, error)
	GetByID(int) (*models.ArticleCopy, error)
	AddArticle(map[string]interface{}) error
	Exists(map[string]interface{}) (bool, error)
	GetsList(page, perPage int, where map[string]interface{}) ([]*models.Article, error)
}

type ArticleService struct {
	manager repositories.IArticleManager
}

func (a *ArticleService) GetsList(page, perPage int, where map[string]interface{}) ([]*models.Article, error) {
	return a.manager.GetsList(page, perPage, where)
}

func NewArticleService() IArticleService {
	return &ArticleService{manager: repositories.NewArticleManager()}
}

func (a *ArticleService) Gets() ([]*models.ArticleCopy, error) {
	list, err := a.manager.Gets()
	if err != nil {
		return nil, err
	}
	data := make([]*models.ArticleCopy, 0)
	for _, d := range list {
		data = append(data, &models.ArticleCopy{Article: *d})
	}
	return data, nil
}

func (a *ArticleService) GetByID(i int) (*models.ArticleCopy, error) {
	articleTag := "article_" + strconv.Itoa(i)
	res, err := gredis.Get(articleTag)
	if err != nil {
		blogger.Error("get article cache error:", err)
	} else {
		var cacheArticle *models.ArticleCopy
		_ = json.Unmarshal(res, &cacheArticle)
		return cacheArticle, nil
	}
	article, err := a.manager.GetsByID(i)
	if err != nil {
		return nil, err
	}
	_ = gredis.Set(articleTag, article, 3600)
	return &models.ArticleCopy{Article: *article}, nil
}

func (a *ArticleService) AddArticle(data map[string]interface{}) error {
	return a.manager.Add(data)
}

func (a *ArticleService) Exists(query map[string]interface{}) (bool, error) {
	return a.manager.Exists(query)
}
