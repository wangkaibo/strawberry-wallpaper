package services

import (
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"log"
)

type StatisticService interface {
	AddVisitLog(ip string, platform string, registerTime string) error
}

type statisticService struct {
	dao *dao.StatisticDao
}

func NewStatisticService() StatisticService {
	return &statisticService{
		dao: dao.NewStatisticDao(db.GetInstanceMysql()),
	}
}

func (s *statisticService) AddVisitLog(ip string, platform string, registerTime string) error {
	err := s.dao.Create(ip, platform, registerTime)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}