package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var ReturnVideos []Video

	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		//VideoList: DemoVideos,
		VideoList: ReturnVideos,
		NextTime:  time.Now().Unix(),
	})
}
