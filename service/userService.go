package service

import "github.com/RaymondCode/simple-demo/dao"

type UserForService struct {
	dao.User
	IsFollow bool `json:"is_follow,omitempty"`
}
