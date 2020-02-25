package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
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
	Version string `json:"version"`
	Username string `json:"username"`
}

func (c *StatisticController) Register(ctx *gin.Context) {
	var registerReq RegisterReq
	err := ctx.ShouldBindJSON(&registerReq)
	if err != nil {
		c.error(ctx,400,err.Error(), gin.H{})
		return
	}
	if registerReq.Uid == "" || registerReq.Version == "" {
		c.error(ctx,400,"参数错误", gin.H{})
		return
	}
	ua := ctx.Request.UserAgent()
	userAgent := user_agent.New(ua)
	user := &models.User{
		Uid: registerReq.Uid,
		Platform: userAgent.Platform(),
		PlatformVersion: userAgent.OSInfo().Version,
		Version: registerReq.Version,
		Username: registerReq.Username,
		RegisterTime: time2.Now(),
		ActiveTime: time2.Now(),
		ActiveDate: time2.Now().Format("2006/01/02"),
		RegisterDate: time2.Now().Format("2006/01/02"),
		Ip: ctx.ClientIP(),
		Ua: ua,
	}
	err = c.StatisticService.Register(user)
	if err != nil {
		c.error(ctx,400, err.Error(), gin.H{})
		return
	}
	c.success(ctx, user)
}

func (c *StatisticController) Active(ctx *gin.Context) {
	activeReq := new(struct{
		Uid string
	})
	raw,_ := ctx.GetRawData()
	err := json.Unmarshal(raw, activeReq)
	if err != nil {
		c.error(ctx,400,"不合法的json", gin.H{})
		return
	}
	err = c.StatisticService.Active(activeReq.Uid)
	if err != nil {
		c.error(ctx, 400, err.Error(), gin.H{})
	} else {
		c.success(ctx, gin.H{})
	}
}


func (c *StatisticController) Index(ctx *gin.Context) {
	startDate := ctx.DefaultQuery("end_date",time2.Now().AddDate(0, 0, -30).Format("2006/01/02"))
	endDate := ctx.DefaultQuery("start_date",time2.Now().Format("2006/01/02"))
	statistic, err := c.StatisticService.GetStatistic(startDate, endDate)
	if err != nil {
		c.error(ctx, 400, err.Error(), gin.H{})
	} else {
		c.success(ctx, statistic)
	}
}
