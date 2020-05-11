package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)


type BaseModel struct {
	gorm.Model
}

func NewDB() *gorm.DB {
	db, err := gorm.Open(viper.GetString("database.driver"),fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.db"),
	))

	if err != nil {
		log.Fatalf("models.NewDB error:%v",err)
	}

	//db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}
