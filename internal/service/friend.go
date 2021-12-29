package service

import (
	"blog/internal/model/result"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetFriend(c *gin.Context) {
	friends, err := s.dao.GetFriendList()
	if err != nil {
		result.Fail(c, result.Error)
		return
	}
	result.Ok(c, friends)
}
