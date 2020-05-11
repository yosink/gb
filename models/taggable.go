package models

// Taggables 标签关联表
type Taggable struct {
	TagID        int64  `gorm:"index;column:tag_id;type:bigint(20) unsigned;not null" json:"tag_id"`           // 标签id
	TaggableID   int64  `gorm:"index;column:taggable_id;type:bigint(20) unsigned;not null" json:"taggable_id"` // 关联对象id
	TaggableType string `gorm:"index;column:taggable_type;type:varchar(191);not null" json:"taggable_type"`    // 关联对象的类型
}
