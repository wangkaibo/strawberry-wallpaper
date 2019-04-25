package db

import (
	"database/sql"
	"sync"
)

var (
	mysql *sql.DB
	lock sync.Mutex
)

func InitMySql() {
	var err interface {Error() string}
	mysql, _ = sql.Open("mysql", "root:123456@/strawberry?charset=utf8")
	if err != nil {
		panic(err)
	}
	return
}

func initMySql() {
	var err interface {Error() string}
	mysql, _ = sql.Open("mysql", "root:123456@/strawberry?charset=utf8")
	if err != nil {
		panic(err)
	}
	return
}
// 没有主从，直接返回单实例
func GetInstanceMysql() *sql.DB {
	if mysql != nil {
		return mysql
	}
	lock.Lock()
	defer lock.Unlock()
	if mysql != nil {
		return mysql
	}
	initMySql()
	return mysql
}

func GetMySql() (*sql.DB) {
	return mysql
}