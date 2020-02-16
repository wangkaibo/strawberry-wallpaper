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

func (dao *NoticeDao) Get(isTest int) (bool, *models.Notice) {
	notice := models.Notice{
		IsTest: isTest,
	}
	has,_ := dao.engine.Where("expire_at>?", time.Now()).Get(&notice)

	return has, &notice
}
