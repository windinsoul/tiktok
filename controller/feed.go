package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []service.VideoForService `json:"video_list,omitempty"`
	NextTime  int64                     `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var ReturnVideos []service.VideoForService
	ReturnVideos, nextTime, err := service.VideoServiceImpl.Feed(time.Now(), int64(1))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		//VideoList: DemoVideos,
		VideoList: ReturnVideos,
		NextTime:  time.Now().Unix(),
	})
}
