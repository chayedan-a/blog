package dao

import (
	"blog/config"
	"blog/internal/model"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func New() *Dao {
	var dao Dao
	data := config.Config()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", data.Database.Redis.Addr, data.Database.Redis.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//rdb.Do("auth" , data.Database.Redis.Password) // 验证密码
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	dao.Redis = rdb

	m := data.Database.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Addr, m.Port, m.Name)
	dao.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = dao.DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Category{}, &model.Friend{}) // 自动创建数据库
	if err != nil {
		panic(err)
	}
	return &dao
}
