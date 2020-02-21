package services

import (
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"strawberry-wallpaper/models"
	"time"
)

type NoticeService interface {
	GetNotices() ([]models.Notice, error)
	GetNoticeList() ([]models.Notice, error)
	DeleteNotice(id string) (bool, error)
	ChangeStatus(id int, isPublish int) error
	AddNotice(content string, expireTime time.Time, publishTime time.Time) error
}

type noticeService struct {
	noticeDao *dao.NoticeDao
}

func NewNoticeService() NoticeService {
	return &noticeService{
		noticeDao: dao.NewNoticeDao(db.GetInstanceMysql()),
	}
}

func (s *noticeService) GetNotices() ([]models.Notice, error) {
	notices, err := s.noticeDao.GetNotices()
	return notices, err
}

func (s *noticeService) GetNoticeList() ([]models.Notice, error) {
	notices, err := s.noticeDao.GetNoticeList()
	return notices, err
}

func (s *noticeService) DeleteNotice(id string) (bool, error) {
	success, err := s.noticeDao.DeleteNotice(id)
	return success, err
}

func (s *noticeService) ChangeStatus(id int, isPublish int) (error) {
	err := s.noticeDao.ChangeStatus(id, isPublish)
	return err
}

func (s *noticeService) AddNotice(content string, publishTime time.Time, expireTime time.Time) (error) {
	notice := models.Notice{
		Content: content,
		IsPublish: 0,
		Time: models.UnixTime(publishTime),
		ExpireAt: models.UnixTime(expireTime),
		CreateAt: models.UnixTime(time.Now()),
	}
	err := s.noticeDao.AddNotice(notice)
	return err
}