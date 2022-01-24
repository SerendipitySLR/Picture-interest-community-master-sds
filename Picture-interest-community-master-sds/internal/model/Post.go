package model

import (
	"time"
)

type Post struct {
	PostId           int `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
	PublisherId      int
	PhotoNumber      int
	Content          string
	CommentNumber    int
	ForwardNumber    int
	LikeNumber       int
	CollectionNumber int
	PhotoPathUrl     string
	Location         string
}

func (Post) TableName() string {
	return "post"
}
