package router

import (
	"blog/config"
	"blog/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(s *service.Service) {
	r := gin.Default()
	r.GET("/friends", s.GetFriend)
	r.GET("/sms/:phone", s.Sms)
	r.POST("/upload", s.Upload)
	panic(r.Run(config.Config().Server.Port))
}
