package dao

import (
	"github.com/go-xorm/xorm"
	"strawberry-wallpaper/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (dao *UserDao) Create(user *models.User) error {
	_, err := dao.engine.Insert(user)
	return err
}

func (dao *UserDao) FindByUid(uid string) (bool, *models.User) {
	user := &models.User{
		Uid: uid,
	}
	has, _ := dao.engine.Get(user)
	if (!has) {
		user = nil
	}
	return has, user
}

func (dao *UserDao) Update(user *models.User) (int64, error) {
	affectRows, err := dao.engine.Id(user.Id).Update(user)
	return affectRows, err
}
