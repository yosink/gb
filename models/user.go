package models

// Users 后台用户表
type User struct {
	BaseModel
	Phone         string `gorm:"unique;column:phone;type:varchar(20);not null" json:"phone"`
	Nickname      string `gorm:"column:nickname;type:varchar(191)" json:"nickname"`
	Avatar        string `gorm:"column:avatar;type:text" json:"avatar"`
	State         uint8  `gorm:"column:state;type:tinyint(4) unsigned;not null" json:"state"` // 0禁用 1正常
	IsAdmin       int8   `gorm:"column:is_admin;type:tinyint(1);not null" json:"is_admin"`
	Password      string `gorm:"column:password;type:varchar(191);not null" json:"password"`
	Description   string `gorm:"column:description;type:varchar(191)" json:"description"`
	RememberToken string `gorm:"column:remember_token;type:varchar(100)" json:"remember_token"`
}
