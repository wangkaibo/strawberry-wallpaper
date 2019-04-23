package db

import "database/sql"

var mysql *sql.DB

func InitMySql() {
	var err interface {Error() string}
	mysql, _ = sql.Open("mysql", "root:123456@/strawberry?charset=utf8")
	if err != nil {
		panic(err)
	}
	return
}

func GetMySql() (*sql.DB) {
	return mysql
}