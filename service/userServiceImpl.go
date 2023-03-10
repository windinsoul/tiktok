package service

import (
	"errors"
	"log"

	"github.com/RaymondCode/simple-demo/dao"
)

type UserServiceImpl struct {
}

func UserIfExist(UserName string) (bool, error) {
	//调用dao层，去查username是否存在
	//存在的话返回true
	exist, err := dao.UserIfExist(UserName)
	if err != nil {
		log.Println("UserIfExist query failed")
		return false, err
	}
	return exist, nil
}

func CreateUser(user dao.User) (bool, error) {
	//调用dao层创建新user
	IfCreated := dao.InsertUser(&user)
	if IfCreated {
		return true, nil
	}
	return false, errors.New("created failed")

}
