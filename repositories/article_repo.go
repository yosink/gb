package repositories

import (
	"blog/models"

	"github.com/jinzhu/gorm"
)

type IArticleManager interface {
	Gets() ([]*models.Article, error)
	GetsByID(int) (*models.Article, error)
	Add(map[string]interface{}) error
	Exists(map[string]interface{}) (bool, error)
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

func (m *ArticleManager) Add(data map[string]interface{}) error {
	article := models.Article{
		CategoryID:      data["cid"].(uint),
		UserID:          data["user_id"].(uint),
		Slug:            data["slug"].(string),
		Title:           data["title"].(string),
		Subtitle:        data["subtitle"].(string),
		Content:         data["content"].(string),
		PageImage:       data["page_image"].(string),
		MetaDescription: data["meta_description"].(string),
		Recommend:       data["recommend"].(*uint8),
		Sort:            data["sort"].(int),
		ViewCount:       data["view_count"].(int),
	}
	return m.DB.Create(&article).Error
}

func (m *ArticleManager) Exists(query map[string]interface{}) (bool, error) {
	var article models.Article
	err := m.DB.Select("id").Where(query).First(&article).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return false, err
	}
	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}
