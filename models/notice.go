package models

import (
	"time"
)

type Notice struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Time     time.Time `json:"time" xorm:"TIMESTAMP"`
	Content  string    `json:"content" xorm:"VARCHAR(1000)"`
	IsTest   int       `json:"is_test" xorm:"not null default 0 TINYINT(1)"`
	ExpireAt time.Time `json:"expire_at" xorm:"TIMESTAMP"`
	CreateAt time.Time `json:"create_at" xorm:"TIMESTAMP"`
}
