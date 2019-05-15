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
	GetStatistic() (map[string]interface{}, error)
}

type statisticService struct {
	userDao *dao.UserDao
	activeDao *dao.ActiveDao
}

func NewStatisticService() StatisticService {
	return &statisticService{
		userDao: dao.NewUserDao(db.GetInstanceMysql()),
		activeDao: dao.NewActiveDao(db.GetInstanceMysql()),
	}
}

func (s *statisticService) Register(user *models.User) error {
	has, _ := s.userDao.FindByUid(user.Uid)
	if has {
		return errors.New("该UID已经注册")
	}
	err := s.userDao.Create(user)
	active := &models.Active{
		Uid: user.Uid,
		ActiveDate: user.ActiveDate,
		ActiveTime: user.ActiveTime,
	}
	_ = s.activeDao.Create(active)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *statisticService) Active(uid string) error {
	has, user := s.userDao.FindByUid(uid)
	if has {
		user.ActiveTime = time.Now()
		user.ActiveDate = time.Now().Format("2006/01/02")
	} else {
		return errors.New("传入的UID不存在")
	}
	_, err := s.userDao.Update(user)
	if err != nil {
		return err
	}
	active := &models.Active{
		Uid: uid,
		ActiveDate: user.ActiveDate,
	}
	_ = s.activeDao.InertOrUpdate(active)
	return nil
}

func (s *statisticService) GetStatistic() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	return data ,nil
}