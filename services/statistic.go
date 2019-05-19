package services

import (
	"errors"
	"log"
	"strawberry-wallpaper/dao"
	"strawberry-wallpaper/db"
	"strawberry-wallpaper/models"
	"time"
)

type StatisticService interface {
	Register(*models.User) error
	Active(string) error
	GetStatistic(startDate string, endDate string) (map[string]interface{}, error)
}

type statisticService struct {
	userDao *dao.UserDao
	activeDao *dao.ActiveDao
}

func NewStatisticService() StatisticService {
	return &statisticService{
		userDao: dao.NewUserDao(db.GetInstanceMysql()),
		activeDao: dao.NewActiveDao(db.GetInstanceMysql()),
	}
}

func (s *statisticService) Register(user *models.User) error {
	has, _ := s.userDao.FindByUid(user.Uid)
	if has {
		return errors.New("该UID已经注册")
	}
	err := s.userDao.Create(user)
	active := &models.Active{
		Uid: user.Uid,
		ActiveDate: user.ActiveDate,
		ActiveTime: user.ActiveTime,
	}
	_ = s.activeDao.Create(active)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *statisticService) Active(uid string) error {
	has, user := s.userDao.FindByUid(uid)
	if has {
		user.ActiveTime = time.Now()
		user.ActiveDate = time.Now().Format("2006/01/02")
	} else {
		return errors.New("传入的UID不存在")
	}
	_, err := s.userDao.Update(user)
	if err != nil {
		return err
	}
	active := &models.Active{
		Uid: uid,
		ActiveDate: user.ActiveDate,
	}
	_ = s.activeDao.InertOrUpdate(active)
	return nil
}

func (s *statisticService) GetStatistic(startDate string, endDate string) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	registerStatistic,_ := s.userDao.GetUserByDate(startDate, endDate)
	activeStatistic,_ := s.activeDao.GetActiveByDate(startDate, endDate)
	currentTime,_ := time.Parse("2006/01/02", startDate);
	endTime, _ := time.Parse("2006/01/02", endDate);
	registerMap := GetDateStatMap(registerStatistic)
	activeMap := GetDateStatMap(activeStatistic)
	regData := []models.DateStat{}
	activeData := []models.DateStat{}
	var currentDate string
	for !currentTime.AddDate(0, 0, 1).After(endTime) {
		currentTime = currentTime.AddDate(0,0,1)
		currentDate = currentTime.Format("2006/01/02")
		emptyStat := models.DateStat{
			Date: currentDate,
			Count: 0,
		}
		if stat ,ok := registerMap[currentDate]; ok {
			regData = append(regData, stat)
		} else {
			regData = append(regData, emptyStat)
		}
		if stat, ok := activeMap[currentDate]; ok {
			activeData = append(activeData, stat)
		} else {
			activeData = append(activeData, emptyStat)
		}
	}
	platformStatistic,_ := s.userDao.GetPlatformStat()
	total,_ := s.userDao.TotalUserNum()
	activeNum,_ := s.userDao.ActiveNum()
	data["register"] = regData
	data["active"] = activeData
	data["platform"] = platformStatistic
	data["total_num"] = total
	data["active_num"] = activeNum
	return data ,nil
}

func GetDateStatMap(s []models.DateStat) map[string]models.DateStat {
	res := map[string]models.DateStat{}
	for _, stat := range(s) {
		res[stat.Date] = stat
	}
	return res
}