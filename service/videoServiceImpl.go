package service

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"
	"log"
	"time"
)

type VideoServiceImpl struct {
}

func (videoService VideoServiceImpl) Feed(LastTime time.Time, userId int64) ([]VideoForService, time.Time, error) {
	videos, err := dao.GetVideoByLastTime(LastTime)
	returnVideos := make([]VideoForService, config.VideoNum)
	if err != nil {
		log.Printf("Failed---->dao.GetVideoByLastTime(LastTime)! : %v", err)
		return nil, time.Time{}, err
	}
	videoService.VideoDaoToService(&videos, &returnVideos, userId)
	return returnVideos, videos[len(videos)-1].PublishTime, nil
}

func (videoService *VideoServiceImpl) VideoDaoToService(daoVideo *[]dao.Video, serviceVideo *[]VideoForService, userId int64) {
	for _, video := range *daoVideo {
		x := VideoForService{
			Video:           video,
			IsWatchedUserId: userId,
			IsFavorite:      false, //之后改
		}
		*serviceVideo = append(*serviceVideo, x)
	}
}
