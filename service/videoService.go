package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"time"
)

type VideoForService struct {
	dao.Video
	Author          UserForService
	IsWatchedUserId int64 `json:"is_watched_user_id"`
	IsFavorite      bool  `json:"is_favorite,omitempty"`
}

type VideoService interface {
	// Feed 返回视频数组，视频数组里最近的时间
	Feed(lastTime time.Time, userId int64) ([]VideoForService, time.Time, error)
}
