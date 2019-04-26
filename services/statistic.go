package services

import (
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"log"
	"strawberry-wallpaper/models"
)

type StatisticService interface {
	Register(*models.User) error
}

type statisticService struct {
	dao *dao.StatisticDao
}

func NewStatisticService() StatisticService {
	return &statisticService{
		dao: dao.NewStatisticDao(db.GetInstanceMysql()),
	}
}

func (s *statisticService) Register(user *models.User) (error) {
	err := s.dao.Create(user)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}