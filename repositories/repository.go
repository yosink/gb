package repositories

import (
	"context"
	"github.com/jinzhu/gorm"
)

type BaseMgr struct {
	*gorm.DB
	ctx       *context.Context
	isRelated bool
}

// SetCtx set context
func (obj *BaseMgr) SetCtx(c *context.Context) {
	obj.ctx = c
}

// GetDB get gorm.DB info
func (obj *BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// GetIsRelated Query foreign key Association.获取是否查询外键关联(gorm.Related)
func (obj *BaseMgr) GetIsRelated() bool {
	return obj.isRelated
}

// SetIsRelated Query foreign key Association.设置是否查询外键关联(gorm.Related)
func (obj *BaseMgr) SetIsRelated(b bool) {
	obj.isRelated = b
}

type options struct {
	query map[string]interface{}
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options)  {
	f(o)
}