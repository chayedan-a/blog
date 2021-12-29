package model

import "time"

type Article struct {
	ID         uint      `gorm:"column:Id;primaryKey"`
	UserId     uint      `gorm:"column:user_id;not null"`
	CategoryId uint      `gorm:"column:category_id;not null"`
	Title      string    `gorm:"column:title;type:varchar(20);not null"`
	Content    string    `gorm:"column:content;type:text;not null"`
	Top        uint      `gorm:"column:top;default:1"`
	View       uint      `gorm:"column:view"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"` // 更新时间
}

func (Article) TableName() string {
	return "t_article"
}
