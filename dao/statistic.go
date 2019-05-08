package dao

import (
	"github.com/go-xorm/xorm"
	"strawberry-wallpaper/models"
)

const (
	_insertLog = "INSERT INTO `users` (`ip`, `os_platform`, `register_time`)" +
		"VALUES (?, ?, ?);"
)

type StatisticDao struct {
	engine *xorm.Engine
}

func NewStatisticDao(engine *xorm.Engine) *StatisticDao {
	return &StatisticDao{
		engine: engine,
	}
}

func (dao *StatisticDao) Create(user *models.User) error {
	_, err := dao.engine.Insert(user)
	return err
}

func (dao *StatisticDao) FindByUid(uid string) (bool, *models.User) {
	user := &models.User{
		Uid: uid,
	}
	has, _ := dao.engine.Get(user)
	if (!has) {
		user = nil
	}
	return has, user
}

func (dao *StatisticDao) Update(user *models.User) (int64, error) {
	affectRows, err := dao.engine.Update(user)
	return affectRows, err
}
