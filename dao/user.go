package dao

import (
	"github.com/go-xorm/xorm"
	"strawberry-wallpaper/models"
	"strconv"
	"time"
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

func (dao *UserDao) GetPlatformStat() ([]map[string]string, error) {
	platformStat := make([]map[string]string, 0)
	platformStat, err := dao.engine.Table("user").Select("platform, count(*) as count").
		GroupBy("platform").QueryString()
	return platformStat, err
}

func (dao *UserDao) TotalUserNum() (int, error) {
	res := make([]map[string]string, 1)
	res, err := dao.engine.Table("user").Select("count(distinct uid) as num").
		QueryString()
	total := 0
	if len(res) > 0 {
		total,_ = strconv.Atoi(res[0]["num"])
	}
	return total, err
}

func (dao *UserDao) ActiveNum() (int, error) {
	res := make([]map[string]string, 1)
	dao.engine.ShowSQL(true)
	res, err := dao.engine.Table("user").Select("count(distinct uid) as num").
		Where("active_date>=?", time.Now().AddDate(0,0, -1).Format("2006/01/02")).
		QueryString()
	activeNum := 0
	if len(res) > 0 {
		activeNum,_ = strconv.Atoi(res[0]["num"])
	}
	return activeNum, err
}
