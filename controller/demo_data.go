package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"time"
)

var DemoVideos = []dao.Video{
	{
		Id:            1,
		AuthorId:      0,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		PublishTime:   time.Time{},
		FavoriteCount: 0,
		CommentCount:  0,
	},
	{
		Id:            2,
		AuthorId:      0,
		PlayUrl:       "https://media.w3.org/2010/05/sintel/trailer.mp4",
		CoverUrl:      "https://ink-blog-1-1308621003.cos.ap-guangzhou.myqcloud.com/img/%E2%9D%89-71348257%20(1).png",
		PublishTime:   time.Time{},
		FavoriteCount: 0,
		CommentCount:  0,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
