package dao

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitDb() {
	Db, err := gorm.Open("mysql", "root:Yyw20011230@(1.12.60.132:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.SingularTable(true) //单数表名
	//defer db.Close()

	//自动迁移
	Db.AutoMigrate(&Video{}, &service.VideoForService{})
}
