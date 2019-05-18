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

func (dao *UserDao) GetUserByDate(startDate string, endDate string) ([]map[string]string, error) {
	userStatistic := make([]map[string]string, 0)
	userStatistic, err := dao.engine.Table("user").Select("register_date, count(*) as count").
		Where("register_date>=? AND register_date<=?", startDate, endDate).
		GroupBy("register_date").QueryString()
	return userStatistic, err
}
