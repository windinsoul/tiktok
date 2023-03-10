package service

import "github.com/RaymondCode/simple-demo/dao"

type UserForService struct {
	dao.User
	IsFollow bool `json:"is_follow,omitempty"`
}
type UserService interface {
	UserIfExist(UserName string) (bool, error)
	CreateUser(user dao.User) (bool, error)
}
