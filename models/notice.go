package models

import (
	"strconv"
	"time"
)

type UnixTime time.Time

// MarshalJSON implements json.Marshaler.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	origin := time.Time(t)
	return []byte(strconv.FormatInt(origin.UnixNano()/1000000, 10)), nil
}

type Notice struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Time     UnixTime `json:"time" xorm:"TIMESTAMP"`
	Content  string    `json:"content" xorm:"VARCHAR(1000)"`
	Status   int       `json:"status" xorm:"not null default 0 TINYINT(1)"`
	ExpireAt UnixTime `json:"expire_at" xorm:"TIMESTAMP"`
	CreateAt UnixTime `json:"create_at" xorm:"TIMESTAMP"`
}