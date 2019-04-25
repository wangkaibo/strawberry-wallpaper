package dao

import (
	"database/sql"
)

const (
	_insertLog = "INSERT INTO `visit_logs` (`ip`, `os_platform`, `register_time`)" +
		"VALUES (?, ?, ?);"
)

type StatisticDao struct {
	db *sql.DB
}

func NewStatisticDao(db *sql.DB) *StatisticDao {
	return &StatisticDao{
		db: db,
	}
}

func (dao *StatisticDao) Create(ip string, platform string, registerTime string) error {
	stmt, err := dao.db.Prepare(_insertLog)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(ip, platform, registerTime)
	return err
}
