package models

import (
	"time"
)

type Active struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Uid        string    `json:"uid" xorm:"not null default '' index VARCHAR(255)"`
	ActiveTime time.Time `json:"active_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ActiveDate string    `json:"active_date" xorm:"not null default '' CHAR(20)"`
}
