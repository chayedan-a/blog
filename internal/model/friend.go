package model

import "time"

type Friend struct {
	ID         uint      `gorm:"column:Id;primaryKey"`
	Url        string    `gorm:"not null;size:128"`
	Name       string    `gorm:"not null;size:32"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"` // 更新时间
}

func (Friend) TableName() string {
	return "t_friend"
}
