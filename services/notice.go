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
	PublishNotice(id int, isPublish int) error
	AddNotice(content string, expireTime time.Time) error
	UpdateNotice(id int, content string, expireTime time.Time) error
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

func (s *noticeService) PublishNotice(id int, isPublish int) (error) {
	notice := models.Notice{
		Id: id,
		IsPublish: isPublish,
	}
	if isPublish == 1 {
		notice.Time = models.UnixTime(time.Now())
	}
	err := s.noticeDao.PublishNotice(notice)
	return err
}

func (s *noticeService) AddNotice(content string, expireTime time.Time) (error) {
	notice := models.Notice{
		Content: content,
		IsPublish: 0,
		ExpireAt: models.UnixTime(expireTime),
		CreateAt: models.UnixTime(time.Now()),
	}
	err := s.noticeDao.AddNotice(notice)
	return err
}

func (s *noticeService) UpdateNotice(id int,content string, expireTime time.Time) (error) {
	notice := models.Notice{
		Id: id,
		Content: content,
		ExpireAt: models.UnixTime(expireTime),
	}
	err := s.noticeDao.UpdateNotice(notice)
	return err
}