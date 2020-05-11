package models

import "time"

// Visitors 内容的浏览记录表
type Visitor struct {
	ID        int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`
	ArticleID int64     `gorm:"index;column:article_id;type:bigint(20) unsigned;not null" json:"article_id"`
	Articles  Article  `gorm:"association_foreignkey:article_id;foreignkey:id" json:"articles_list"` // 内容表
	IP        string    `gorm:"column:ip;type:varchar(32);not null" json:"ip"`
	Country   string    `gorm:"column:country;type:varchar(191)" json:"country"`
	Clicks    int       `gorm:"column:clicks;type:int(10) unsigned;not null" json:"clicks"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}
