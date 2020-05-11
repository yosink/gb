package repositories

import (
	"blog/models"
)

type IArticleManager interface {
	Gets() ([]*models.Article, error)
	GetsByID(int) (models.Article, error)
}

type ArticleManager struct {
	*BaseMgr
}

func NewArticleManager() *ArticleManager {
	return &ArticleManager{
		BaseMgr: &BaseMgr{
			DB:        models.NewDB(),
		},
	}
}

func (m *ArticleManager) GetsByID(id int) (result models.Article, err error) {
	err = m.DB.Where("id = ?", id).Find(&result).Error
	return
}

func (m *ArticleManager) Gets() (result []*models.Article, err error) {
	err = m.DB.Find(&result).Error
	return
}
