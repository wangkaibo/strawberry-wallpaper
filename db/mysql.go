package db

import (
	"os"
	"sync"
	"github.com/go-xorm/xorm"
)

var (
	mysql *xorm.Engine
	lock sync.Mutex
)

func initMySql() {
	var err error
	dataSourceName := os.Getenv("MYSQL_USERNAME")
	dataSourceName += ":" + os.Getenv("MYSQL_PASSWORD")
	dataSourceName += "@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	dataSourceName +=  "/strawberry?charset=utf8"
	mysql, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return
}
// 没有主从，直接返回单实例
func GetInstanceMysql() *xorm.Engine {
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