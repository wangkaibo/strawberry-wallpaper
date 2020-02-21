package dao

import (
	"github.com/go-xorm/xorm"
	"strawberry-wallpaper/models"
	"strconv"
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

func (dao *NoticeDao) GetNotices() ([]models.Notice, error) {
	notices := make([]models.Notice, 0)
	err := dao.engine.Where("expire_at>?", time.Now()).Where("is_publish=?", 1).Limit(50).Find(&notices)
	return notices, err
}

func (dao *NoticeDao) GetNoticeList() ([]models.Notice, error) {
	notices := make([]models.Notice, 0)
	err := dao.engine.Table("notice_test").OrderBy("create_at DESC").Limit(100).Find(&notices)
	return notices, err
}

func (dao *NoticeDao) DeleteNotice(id string) (bool, error) {
	idInt, _ := strconv.Atoi(id)
	notice := &models.Notice{
		Id: idInt,
	}
	_, err := dao.engine.Table("notice_test").Delete(notice)
	res := false
	if err == nil {
		res = true
	}
	return res, err
}

func (dao *NoticeDao) ChangeStatus(id int, isPublish int) (error) {
	notice := models.Notice{
		Id: id,
		IsPublish: isPublish,
	}
	_, err := dao.engine.Table("notice_test").Cols("is_publish").Update(&notice)

	return err
}
func (dao *NoticeDao) AddNotice(notice models.Notice) (error) {
	_, err := dao.engine.Table("notice_test").Insert(&notice)

	return err
}