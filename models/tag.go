package models

// Tags 标签表
type Tag struct {
	BaseModel
	Name            string `gorm:"index;column:name;type:varchar(191);not null" json:"name"`                   // 标签名称
	Img             string `gorm:"column:img;type:varchar(255)" json:"img"`                                    // 标签图片
	MetaDescription string `gorm:"column:meta_description;type:varchar(191);not null" json:"meta_description"` // 描述
	State           int8   `gorm:"column:state;type:tinyint(3)" json:"state"`                                  // 状态
	Sort            int    `gorm:"column:sort;type:int(11)" json:"sort"`                                       // 排序
}
