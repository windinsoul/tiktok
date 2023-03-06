package dao

import (
	"github.com/RaymondCode/simple-demo/config"
	"time"
)

type Video struct {
	Id            int64 `json:"id,omitempty"`
	AuthorId      int64
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	PublishTime   time.Time
	Author        User  `json:"author"`
	FavoriteCount int64 `json:"favorite_count,omitempty"`
	CommentCount  int64 `json:"comment_count,omitempty"`
}

func AddVideoToDb(video []Video, Time time.Time) error {
	for _, x := range video {
		if result := Db.Create(&x); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func GetVideoByLastTime(lastTime time.Time) ([]Video, error) {
	ret := make([]Video, config.VideoNum)
	if res := Db.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(config.VideoNum).Find(&ret); res.Error != nil {
		return ret, res.Error
	}
	return ret, nil
}
