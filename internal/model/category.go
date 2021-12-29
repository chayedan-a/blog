package model

import (
	"time"
)

type Category struct {
	ID         uint      `gorm:"column:Id;primaryKey"`
	Name       string    `gorm:"size:16;not null"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"` // 更新时间
}

func (Category) TableName() string {
	return "t_category"
}
