package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strawberry-wallpaper/models"
	"strawberry-wallpaper/services"
	time2 "time"
)

type StatisticController struct {
	BaseController
	StatisticService services.StatisticService
}

type RegisterReq struct {
	Uid string `json:"uid"`
	Platform string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
	Version string `json:"version"`
	Username string `json:"username"`
}

func (c *StatisticController) Register(ctx *gin.Context) {
	registerReq := &RegisterReq{}
	raw,_ := ctx.GetRawData()
	json.Unmarshal(raw, registerReq)
	user := &models.User{
		Uid: registerReq.Uid,
		Platform: registerReq.Platform,
		PlatformVersion: registerReq.PlatformVersion,
		Version: registerReq.Version,
		Username: registerReq.Username,
		RegisterTime: time2.Now(),
		ActiveTime: time2.Now(),
		Ip: ctx.ClientIP(),
		Ua: ctx.Request.UserAgent(),
	}
	c.StatisticService.Register(user)
	c.success(ctx, user)
}

func (c *StatisticController) Active(ctx *gin.Context) {
	activeReq := new(struct{
		Uid string
	})
	raw,_ := ctx.GetRawData()
	json.Unmarshal(raw, activeReq)
	err := c.StatisticService.Active(activeReq.Uid)
	if err != nil {
		c.error(ctx, 400, err.Error(), gin.H{})
	} else {
		c.success(ctx, gin.H{})
	}
}
