package dao

import (
	"github.com/RaymondCode/simple-demo/config"
	"log"
	"time"
)

type Video struct {
	Id            int64 `json:"id"`
	AuthorId      int64
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	PublishTime   time.Time
	FavoriteCount int64 `json:"favorite_count"`
	CommentCount  int64 `json:"comment_count"`
}

func AddVideoToDb(video []Video) error {
	for _, x := range video {
		if result := Db.Create(&x); result.Error != nil {
			return result.Error
		}
	}
	log.Printf("AddVideoToDb Successfully!")
	return nil
}

func GetVideoByLastTime(lastTime time.Time) ([]Video, error) {
	ret := make([]Video, config.VideoNum)
	var num int
	if res := Db.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(config.VideoNum).Find(&ret).Count(&num); res.Error != nil {
		return ret, res.Error
	}
	//fmt.Printf("%v\n", ret)
	//fmt.Printf("%d\n", num)
	return ret, nil
}
