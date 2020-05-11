package models

// Links [...]
type Link struct {
	BaseModel
	Name   string `gorm:"index;column:name;type:varchar(191);not null" json:"name"`
	Link   string `gorm:"index;column:link;type:varchar(191);not null" json:"link"`
	Image  string `gorm:"column:image;type:text" json:"image"`
	Status int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`
}
