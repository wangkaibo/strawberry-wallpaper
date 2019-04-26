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
	Service services.StatisticService
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
		Ip: ctx.ClientIP(),
		Ua: ctx.Request.UserAgent(),
	}
	c.Service.Register(user)
	c.success(ctx, user)
}
