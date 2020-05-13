package models

import (
	"blog/api/resources"
	"blog/comm"
	"encoding/json"
)

// Articles 内容表
type Article struct {
	BaseModel
	CategoryID      uint `gorm:"column:category_id;type:bigint(20) unsigned;not null"` // 分类
	Category        *Category
	UserID          uint `gorm:"column:user_id;type:bigint(20) unsigned;not null"` // 作者
	User            *User
	Slug            string     `gorm:"unique;column:slug;type:varchar(191);not null"`           // 短链
	Title           string     `gorm:"column:title;type:varchar(191);not null" `                // 标题
	Subtitle        string     `gorm:"column:subtitle;type:varchar(191);not null" `             // 副标题
	Content         string     `gorm:"column:content;type:text;not null" `                      // 内容
	PageImage       string     `gorm:"column:page_image;type:varchar(191)" `                    // 主图
	MetaDescription string     `gorm:"column:meta_description;type:varchar(191)" `              // seo内容
	Recommend       *uint8     `gorm:"column:recommend;type:tinyint(1);not null" `              // 是否推荐
	Sort            int        `gorm:"column:sort;type:int(4) unsigned;not null" `              // 排序
	State           uint8      `gorm:"column:state;type:tinyint(1);not null" `                  // 0默认草稿 1已发布
	ViewCount       int        `gorm:"index;column:view_count;type:int(10) unsigned;not null" ` // 浏览量
	PublishedAt     comm.XTime `gorm:"column:published_at;type:timestamp" `                     // 发布时间
}

type ArticleCopy struct {
	Article
}

func (a *ArticleCopy) MarshalJSON() ([]byte, error) {
	article := resources.ArticleResource{
		ID:              a.ID,
		Slug:            a.Slug,
		Title:           a.Title,
		SubTitle:        a.Subtitle,
		Content:         a.Content,
		PageImage:       a.PageImage,
		MetaDescription: a.MetaDescription,
		Recommend:       a.Recommend,
		Sort:            a.Sort,
		ViewCount:       a.ViewCount,
		CreatedAt:       a.CreatedAt,
		User:            nil,
	}
	if a.Category != nil {
		article.Category = &resources.CategoryResource{
			ID:          a.CategoryID,
			Name:        a.Category.Name,
			Description: a.Category.Description,
			ImageURL:    a.Category.ImageURL,
			CreatedAt:   a.Category.CreatedAt,
		}
	}
	return json.Marshal(article)
}
