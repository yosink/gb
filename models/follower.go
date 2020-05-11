package models

import "time"

// Followers [...]
type Follower struct {
	ID        int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`
	UserID    int64     `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`
	FollowID  int64     `gorm:"column:follow_id;type:bigint(20) unsigned;not null" json:"follow_id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}
