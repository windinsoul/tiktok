package dao

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitDb() {
	Db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.SingularTable(true) //单数表名
	//defer db.Close()

	//自动迁移
	Db.AutoMigrate(&controller.Video{}, &controller.VideoPublishTime{})
}