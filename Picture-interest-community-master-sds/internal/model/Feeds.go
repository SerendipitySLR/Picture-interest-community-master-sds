package model

import "time"

type Feeds struct {
	UserId    int `gorm:"primary_key"`
	PostId    int `gorm:"primary_key"`
	PostType  int `gorm:"primary_key"`
	SendId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Feeds) TableName() string {
	return "feeds"
}
