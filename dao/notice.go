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

func (dao *NoticeDao) GetNotices(status int) ([]models.Notice, error) {
	notices := make([]models.Notice, 0)
	err := dao.engine.Where("expire_at>?", time.Now()).Where("status=?", status).Limit(50).Find(&notices)
	return notices, err
}
