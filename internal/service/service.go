package service

import (
	"blog/internal/dao"
)

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	srv := Service{
		dao: dao.New(),
	}
	return &srv
}
