package model

import "time"

type Forward struct {
	ForwardId     int `gorm:"primary_key"`
	PostId        int
	UserId        int
	state         int
	Content       string
	CommentNumber int
	LikeNumber    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func (Forward) TableName() string {
	return "forward"
}
