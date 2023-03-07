package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []service.VideoForService `json:"video_list,omitempty"`
	NextTime  int64                     `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	var lastTime time.Time
	if latestTime != "" {
		rbq, _ := strconv.ParseInt(latestTime, 10, 64)
		lastTime = time.Unix(rbq, 0) //时间戳转化为标准时间
	} else {
		lastTime = time.Now()
	}

	var userId int64 = 0 //token还没弄好
	//userId, _ = strconv.ParseInt(c.Query("token"), 10, 64)
	ReturnVideos, NextTime, err := service.Feed(lastTime, userId)
	if err != nil {
		panic(err)
	}

	//存该用户的nextTime为LastTime
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		//VideoList: DemoVideos,
		VideoList: ReturnVideos,
		NextTime:  NextTime.Unix(),
	})
}
