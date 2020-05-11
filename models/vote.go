package models

import (
	"time"
)

// Votes 点赞表
type Vote struct {
	UserID      int64     `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`          // 用户id
	VotableID   int64     `gorm:"column:votable_id;type:bigint(20) unsigned;not null" json:"votable_id"`    // 点赞对象id
	VotableType string    `gorm:"index;column:votable_type;type:varchar(191);not null" json:"votable_type"` // 点赞对象类型
	Type        string    `gorm:"column:type;type:enum('up_vote','down_vote');not null" json:"type"`        // 点赞或鄙视
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}
