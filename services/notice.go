package services

import (
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"strawberry-wallpaper/models"
)

type NoticeService interface {
	GetNotice(int) (bool, *models.Notice)
}

type noticeService struct {
	noticeDao *dao.NoticeDao
}

func NewNoticeService() NoticeService {
	return &noticeService{
		noticeDao: dao.NewNoticeDao(db.GetInstanceMysql()),
	}
}

func (s *noticeService) GetNotice(isTest int) (bool, *models.Notice) {
	has,notice := s.noticeDao.Get(isTest)
	return has, notice
}