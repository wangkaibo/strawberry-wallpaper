package models

import (
	"strawberry/db"
)

const (
	_insertLog = "INSERT INTO `visit_logs` (`ip`, `os_platform`, `register_time`)" +
		"VALUES (?, ?, ?);"
)

func AddVisitLog(ip string, platform string, registerTime string) {
	mysql := db.GetMySql()
	stmt, err := mysql.Prepare(_insertLog)
	if err != nil {
		panic(err)
	}
	stmt.Exec(ip, platform, registerTime)
	return
}