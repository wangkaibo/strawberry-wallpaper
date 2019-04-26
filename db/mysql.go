package db

import (
	"database/sql"
	"os"
	"sync"
)

var (
	mysql *sql.DB
	lock sync.Mutex
)

func initMySql() {
	var err error
	dataSourceName := os.Getenv("MYSQL_USERNAME")
	dataSourceName += ":" + os.Getenv("MYSQL_PASSWORD")
	dataSourceName += "@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	dataSourceName +=  "/strawberry?charset=utf8"
	mysql, err = sql.Open("mysql", dataSourceName)
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