package controllers

import (
	"github.com/gin-gonic/gin"
	"strawberry-wallpaper/utils"
	time2 "time"
	"strawberry-wallpaper/services"
	"strawberry-wallpaper/models"
)

type StatisticController struct {
	BaseController
	StatisticService services.StatisticService
}

func (c *StatisticController) Register(ctx *gin.Context) {
	platform := utils.GetPlatformByUa(ctx.Request.UserAgent())
	platformVersion := ctx.PostForm("platformVersion")
	userName := ctx.PostForm("username")
	uid := ctx.PostForm("uid")
	user := &models.User{
		Uid: uid,
		Platform: platform,
		PlatformVersion: platformVersion,
		Username: userName,
		RegisterTime: time2.Now(),
		ActiveTime: time2.Now(),
		Ip: ctx.ClientIP(),
		Ua: ctx.Request.UserAgent(),
	}
	c.StatisticService.Register(user)
	c.success(ctx, user)
}

func (c *StatisticController) Active(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	err := c.StatisticService.Active(uid)
	if err != nil {
		c.error(ctx, 400, err.Error(), gin.H{})
	} else {
		c.success(ctx, gin.H{})
	}
}
