package models

import "time"

// PasswordResets [...]
type PasswordReset struct {
	Email     string    `gorm:"index;column:email;type:varchar(191);not null" json:"email"`
	Token     string    `gorm:"index;column:token;type:varchar(191);not null" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
}
