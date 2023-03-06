package dao

import (
	"github.com/RaymondCode/simple-demo/controller"
	"time"
)

func AddVideoToDb(video controller.Video, Time time.Time) error {
	if result := Db.Create(&video); result.Error != nil {
		return result.Error
	}

	videoPublishTime := controller.VideoPublishTime{
		Video:       video,
		PublishTime: Time,
	}
	if result := Db.Create(&videoPublishTime); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetVideoFromDb(last_time time.Time) ([]controller.Video, error) {
	var ret []controller.Video
	Db.Where("id < ?", "31").Find(&ret)
	return ret, nil
}
