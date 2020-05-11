package models

// Categories 内容分类
type Category struct {
	BaseModel
	ParentID    int64  `gorm:"column:parent_id;type:bigint(20) unsigned;not null" json:"parent_id"` // 父级id
	Name        string `gorm:"column:name;type:varchar(191);not null" json:"name"`                  // 分类名称
	Description string `gorm:"column:description;type:varchar(191)" json:"description"`             // 描述
	ImageURL    string `gorm:"column:image_url;type:varchar(191)" json:"image_url"`                 // 分类图片
}
