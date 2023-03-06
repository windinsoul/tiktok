package dao

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
	Work_count      int64  `json:"work___count,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
}
