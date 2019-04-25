package controllers

import (
	"github.com/gin-gonic/gin"
	"strawberry-wallpaper/utils"
	time2 "time"
	"strawberry-wallpaper/services"
)

var (
	service services.StatisticService
)

func init() {
	service = services.NewStatisticService()
}

type StatisticController struct {
	BaseController
}

func (controller *StatisticController) Statistic(c *gin.Context) {
	platform := utils.GetPlatformByUa(c.Request.UserAgent())
	registerTime := time2.Now().Format("2006-01-02 15:04:05")
	service.AddVisitLog(c.Request.RemoteAddr, platform, registerTime)
	controller.success(c, map[string]interface{}{
		"ip": c.Request.RemoteAddr,
		"platform": platform,
		"register_time": registerTime,
	})
}
