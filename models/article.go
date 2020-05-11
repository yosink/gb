package models

import (
	"time"
)

// Articles 内容表
type Article struct {
	BaseModel
	CategoryID      int64 `gorm:"column:category_id;type:bigint(20) unsigned;not null" json:"category_id"` // 分类
	Category        Category
	UserID          int64 `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"user_id"` // 作者
	User            User
	Slug            string    `gorm:"unique;column:slug;type:varchar(191);not null" json:"slug"`                // 短链
	Title           string    `gorm:"column:title;type:varchar(191);not null" json:"title"`                     // 标题
	Subtitle        string    `gorm:"column:subtitle;type:varchar(191);not null" json:"subtitle"`               // 副标题
	Content         string    `gorm:"column:content;type:text;not null" json:"content"`                         // 内容
	PageImage       string    `gorm:"column:page_image;type:varchar(191)" json:"page_image"`                    // 主图
	MetaDescription string    `gorm:"column:meta_description;type:varchar(191)" json:"meta_description"`        // seo内容
	Recommend       int8      `gorm:"column:recommend;type:tinyint(1);not null" json:"recommend"`               // 是否推荐
	Sort            int       `gorm:"column:sort;type:int(4) unsigned;not null" json:"sort"`                    // 排序
	State           int8      `gorm:"column:state;type:tinyint(1);not null" json:"state"`                       // 0默认草稿 1已发布
	ViewCount       int       `gorm:"index;column:view_count;type:int(10) unsigned;not null" json:"view_count"` // 浏览量
	PublishedAt     time.Time `gorm:"column:published_at;type:timestamp" json:"published_at"`                   // 发布时间
}
