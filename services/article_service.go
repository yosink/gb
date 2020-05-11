package services

import (
	"blog/models"
	"blog/repositories"
)

type IArticleService interface {
	Gets() ([]*models.Article, error)
	GetsByID(int) (models.Article, error)
}

type ArticleService struct {
	manager repositories.IArticleManager
}

func NewArticleService() IArticleService {
	return &ArticleService{manager: repositories.NewArticleManager()}
}

func (a ArticleService) Gets() ([]*models.Article, error) {
	return a.manager.Gets()
}

func (a ArticleService) GetsByID(i int) (models.Article, error) {
	return a.manager.GetsByID(i)
}
