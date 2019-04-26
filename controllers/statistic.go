package controllers

import (
	"github.com/gin-gonic/gin"
	"strawberry-wallpaper/utils"
	time2 "time"
	"strawberry-wallpaper/services"
)

type StatisticController struct {
	BaseController
	Service services.StatisticService
}

func (c *StatisticController) Register(ctx *gin.Context) {
	platform := utils.GetPlatformByUa(ctx.Request.UserAgent())
	registerTime := time2.Now().Format("2006-01-02 15:04:05")
	c.Service.AddVisitLog(ctx.Request.RemoteAddr, platform, registerTime)
	c.success(ctx, map[string]interface{}{
		"ip": ctx.Request.RemoteAddr,
		"platform": platform,
		"register_time": registerTime,
	})
}
