package repositories

import (
	"blog/models"
)

type IArticleManager interface {
	Gets() ([]*models.Article, error)
	GetsByID(int) (*models.Article, error)
}

type ArticleManager struct {
	*BaseMgr
}

func NewArticleManager() *ArticleManager {
	return &ArticleManager{
		BaseMgr: &BaseMgr{
			DB: models.NewDB(),
		},
	}
}

func (m *ArticleManager) GetsByID(id int) (*models.Article, error) {
	//err = m.DB.Where("id = ?", id).Find(&result).Error
	//return
	var result models.Article
	d := m.DB.Model(&models.Article{}).Where("id = ?", id)
	d = d.Preload("Category").Preload("User")
	err := d.Find(&result).Error
	//m.DB.Model(&models.Article{}).Preload("User")
	//err = m.DB.Find(&result).Error
	return &result, err
}

func (m *ArticleManager) Gets() (result []*models.Article, err error) {
	err = m.DB.Find(&result).Error
	return
}
