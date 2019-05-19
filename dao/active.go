package dao

import (
	"github.com/go-xorm/xorm"
	"strawberry-wallpaper/models"
	"time"
)

type ActiveDao struct {
	engine *xorm.Engine
}

func NewActiveDao(engine *xorm.Engine) *ActiveDao {
	return &ActiveDao{
		engine: engine,
	}
}

func (dao *ActiveDao) InertOrUpdate(active *models.Active) error {
	exist,_ := dao.engine.Exist(active)
	var err error
	active.ActiveTime = time.Now()
	if exist {
		_, err = dao.engine.Where("uid=? AND active_date=?", active.Uid, active.ActiveDate).Update(active)
	} else {
		_, err = dao.engine.Insert(active)
	}
	return err
}

func (dao *ActiveDao) Create(active *models.Active) error {
	_, err := dao.engine.Insert(active)
	return err
}

func (dao *ActiveDao) GetActiveByDate(startDate string, endDate string) ([]models.DateStat, error) {
	dateStat := make([]models.DateStat, 0)
	err := dao.engine.Table("active").Select("active_date as date, count(*) as count").
		Where("active_date>=? AND active_date<=?", startDate, endDate).
		GroupBy("active_date").Find(&dateStat)
	return dateStat, err
}
