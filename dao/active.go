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
		_, err = dao.engine.Where("uid=?", active.Uid).Update(active)
	} else {
		_, err = dao.engine.Insert(active)
	}
	return err
}

func (dao *ActiveDao) Create(active *models.Active) error {
	_, err := dao.engine.Insert(active)
	return err
}
