package router

import (
	"blog/config"
	"blog/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(s *service.Service) {
	r := gin.Default()
	r.GET("/Friends", s.GetFriend)
	panic(r.Run(config.Config().Server.Port))
}
