package models

// Comments 评论表
type Comment struct {
	BaseModel
	UserID    int    `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"` // 发表评论的用户
	ArticleID int    `gorm:"column:article_id;type:int(10);not null" json:"article_id"`
	Content   string `gorm:"column:content;type:text;not null" json:"content"`   // 评论内容
	Sort      int    `gorm:"column:sort;type:int(4)" json:"sort"`                // 排序
	State     int8   `gorm:"column:state;type:tinyint(1);not null" json:"state"` // 1正常显示 0不显示
}
