package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB

func InitDb() {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.SingularTable(true) //单数表名
	//defer db.Close()

	//自动迁移
	Db.AutoMigrate(&User{}, &Video{})

}

var DemoVideos = []Video{
	{
		Id:            1,
		AuthorId:      0,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		PublishTime:   time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local),
		FavoriteCount: 0,
		CommentCount:  0,
	},
	{
		Id:            2,
		AuthorId:      0,
		PlayUrl:       "https://media.w3.org/2010/05/sintel/trailer.mp4",
		CoverUrl:      "https://ink-blog-1-1308621003.cos.ap-guangzhou.myqcloud.com/img/%E2%9D%89-71348257%20(1).png",
		PublishTime:   time.Date(2020, 2, 1, 1, 1, 1, 1, time.Local),
		FavoriteCount: 0,
		CommentCount:  0,
	},
}
