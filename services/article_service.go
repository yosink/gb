package services

import (
	"blog/models"
	"blog/repositories"
)

type IArticleService interface {
	Gets() ([]*models.ArticleCopy, error)
	GetByID(int) (*models.ArticleCopy, error)
}

type ArticleService struct {
	manager repositories.IArticleManager
}

func NewArticleService() IArticleService {
	return &ArticleService{manager: repositories.NewArticleManager()}
}

func (a ArticleService) Gets() ([]*models.ArticleCopy, error) {
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

func (a ArticleService) GetByID(i int) (*models.ArticleCopy, error) {
	article, err := a.manager.GetsByID(i)
	if err != nil {
		return nil, err
	}
	return &models.ArticleCopy{Article: *article}, nil
}
