package services

import (
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"strawberry-wallpaper/models"
)

type NoticeService interface {
	GetNotices(int) ([]models.Notice, error)
}

type noticeService struct {
	noticeDao *dao.NoticeDao
}

func NewNoticeService() NoticeService {
	return &noticeService{
		noticeDao: dao.NewNoticeDao(db.GetInstanceMysql()),
	}
}

func (s *noticeService) GetNotices(isTest int) ([]models.Notice, error) {
	notices, err := s.noticeDao.GetNotices(isTest)
	return notices, err
}