package services

import (
	"blog/models"
	"blog/repositories"
)

type IArticleService interface {
	Gets() ([]*models.ArticleCopy, error)
	GetByID(int) (*models.ArticleCopy, error)
	AddArticle(map[string]interface{}) error
	Exists(map[string]interface{}) (bool, error)
}

type ArticleService struct {
	manager repositories.IArticleManager
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
	article, err := a.manager.GetsByID(i)
	if err != nil {
		return nil, err
	}
	return &models.ArticleCopy{Article: *article}, nil
}

func (a *ArticleService) AddArticle(data map[string]interface{}) error {
	return a.manager.Add(data)
}

func (a *ArticleService) Exists(query map[string]interface{}) (bool, error) {
	return a.manager.Exists(query)
}
