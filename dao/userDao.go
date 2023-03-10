package dao

import (
	"log"
)

type User struct {
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Password        string `json:"password,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signatrue       string `json:"signatrue,omitempty"`
	TotalFavorited  int64  `json:"total_favorited,omitempty"`
	Work_count      int64  `json:"work_count,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
}

func GetUserById(id int64) (User, error) {
	user := User{}
	ret := Db.Where("id = ?", id).First(&user)
	if ret.Error != nil {
		log.Println(ret.Error)
		return user, ret.Error
	}
	return user, nil
}

func GetUserByName(name string) (User, error) {
	user := User{}
	ret := Db.Where("name = ?", name).First(&user)
	if ret.Error != nil {
		log.Println(ret.Error)
		return user, ret.Error
	}
	return user, nil
}

func InsertUser(user *User) bool {
	if err := Db.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func UserIfExist(username string) (bool, error) {
	ret := Db.Where("name = ?", username).First(&User{})
	if ret.Error != nil {
		if ret.RecordNotFound() {
			return false, nil
		} else {
			log.Printf(ret.Error.Error())
			return false, ret.Error
		}
	}
	return true, nil
}
