package models

// Categories 内容分类
type Category struct {
	BaseModel
	ParentID    int64  `gorm:"column:parent_id;type:bigint(20) unsigned;not null"` // 父级id
	Name        string `gorm:"column:name;type:varchar(191);not null" `            // 分类名称
	Description string `gorm:"column:description;type:varchar(191)" `              // 描述
	ImageURL    string `gorm:"column:image_url;type:varchar(191)" `                // 分类图片
}
