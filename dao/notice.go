package dao

import (
	"github.com/go-xorm/xorm"
	"strawberry-wallpaper/models"
	"time"
)

type NoticeDao struct {
	engine *xorm.Engine
}

func NewNoticeDao(engine *xorm.Engine) *NoticeDao {
	return &NoticeDao{
		engine: engine,
	}
}

func (dao *NoticeDao) GetNotices(isPublish int) ([]models.Notice, error) {
	notices := make([]models.Notice, 0)
	err := dao.engine.Where("expire_at>?", time.Now()).Where("is_publish=?", isPublish).Limit(50).Find(&notices)
	return notices, err
}

func (dao *NoticeDao) GetNoticeList() ([]models.Notice, error) {
	notices := make([]models.Notice, 0)
	err := dao.engine.OrderBy("create_at DESC").Limit(100).Find(&notices)
	return notices, err
}