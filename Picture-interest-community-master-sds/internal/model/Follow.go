package model

import "time"

type Follow struct {
	UserId    int `gorm:"primary_key"`
	FollowId  int
	State     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Follow) TableName() string {
	return "follow"
}
