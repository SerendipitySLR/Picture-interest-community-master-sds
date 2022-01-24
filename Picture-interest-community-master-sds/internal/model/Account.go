package model

import "time"

//用户相关信息
type UserRegister struct {
}

type UserDetails struct {
	UserId     int `gorm:"primary_key"`
	NickName   string
	Sex        bool
	Birthday   string
	Location   string
	Signature  string
	ProfileUrl string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type UserRelatedData struct {
	UserId           int `gorm:"primary_key"`
	FollowsNumber    int
	FansNumber       int
	PostsNumber      int
	CollectionNumber int
	ForwardNumber    int
	ProfileUrl       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func (UserDetails) TableName() string {
	return "user_details"
}
func (UserRelatedData) TableName() string {
	return "user_related_data"
}
