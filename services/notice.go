package services

import (
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"strawberry-wallpaper/models"
)

type NoticeService interface {
	GetNotices(int) ([]models.Notice, error)
	GetNoticeList() ([]models.Notice, error)
}

type noticeService struct {
	noticeDao *dao.NoticeDao
}

func NewNoticeService() NoticeService {
	return &noticeService{
		noticeDao: dao.NewNoticeDao(db.GetInstanceMysql()),
	}
}

func (s *noticeService) GetNotices(isPublish int) ([]models.Notice, error) {
	notices, err := s.noticeDao.GetNotices(isPublish)
	return notices, err
}

func (s *noticeService) GetNoticeList() ([]models.Notice, error) {
	notices, err := s.noticeDao.GetNoticeList()
	return notices, err
}