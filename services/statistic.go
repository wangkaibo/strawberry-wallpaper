package services

import (
	"errors"
	"log"
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"strawberry-wallpaper/models"
	"time"
)

type StatisticService interface {
	Register(*models.User) error
	Active(string) error
}

type statisticService struct {
	dao *dao.StatisticDao
}

func NewStatisticService() StatisticService {
	return &statisticService{
		dao: dao.NewStatisticDao(db.GetInstanceMysql()),
	}
}

func (s *statisticService) Register(user *models.User) error {
	has, _ := s.dao.FindByUid(user.Uid)
	if has {
		return errors.New("该UID已经注册")
	}
	err := s.dao.Create(user)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *statisticService) Active(uid string) error {
	has, user := s.dao.FindByUid(uid)
	if has {
		user.ActiveTime = time.Now()
	} else {
		return errors.New("传入的UID不存在")
	}
	_, err := s.dao.Update(user)
	if err != nil {
		return err
	}
	return nil
}