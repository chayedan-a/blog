package service

import (
	"blog/internal/dao"
	"blog/internal/model/result"
)

type Service struct {
	dao    *dao.Dao
	ErrMap map[uint]string
}

func New() *Service {
	srv := Service{
		dao:    dao.New(),
		ErrMap: result.GetErrMap(),
	}
	return &srv
}
