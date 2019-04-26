package models

import (
	"time"
)

type User struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Platform        string    `json:"platform" xorm:"not null default 'windows' ENUM('osx','windows')"`
	PlatformVersion string    `json:"platform_version" xorm:"not null default '' VARCHAR(255)"`
	Username        string    `json:"username" xorm:"not null default '' VARCHAR(255)"`
	Uid             string    `json:"uid" xorm:"not null default '' VARCHAR(255)"`
	RegisterTime    time.Time `json:"register_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Ip              string    `json:"ip" xorm:"not null VARCHAR(255)"`
	Ua              string    `json:"ua" xorm:"not null default '' VARCHAR(255)"`
}
