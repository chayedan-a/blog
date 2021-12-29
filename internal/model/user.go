package model

import (
	"time"
)

// User 对应t_user表字段
type User struct {
	ID         uint      `gorm:"column:Id;primaryKey"`
	Name       string    `gorm:"type:varchar(20);not null"`
	Password   string    `gorm:"type:varchar(64);not null"`
	Phone      string    `gorm:"type:varchar(11);not null;unique"`
	Icon       string    `gorm:"type:varchar(255)"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"` // 更新时间
}

// TableName 会将User 结构体的表名改为 t_user
func (User) TableName() string {
	return "t_user"
}

// UserInfo 用户基本信息
type UserInfo struct {
	ID    uint
	Name  string
	Icon  string
	Phone string
}

// RegisterReq 注册请求参数
type RegisterReq struct {
	Name     string
	Password string
	Phone    string
	Icon     string
	Code     string
}

// UpdateUserInfoReq 修改用户信息请求参数
type UpdateUserInfoReq struct {
	ID   uint
	Name string
	Icon string
}

// LoginReq 登录请求参数
type LoginReq struct {
	Phone    string
	Password string
}

// LoginRsp 登录返回参数
type LoginRsp struct {
	userInfo UserInfo
	token    string
}
