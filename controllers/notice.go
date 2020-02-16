package controllers

import (
	"github.com/gin-gonic/gin"
	"strawberry-wallpaper/models"
	"strawberry-wallpaper/services"
)

type NoticeController struct {
	BaseController
	NoticeService services.NoticeService
}

func (c *NoticeController) Notice(ctx *gin.Context) {
	paramTest := ctx.Query("is_test")
	isTest := 0
	if paramTest == "1" {
		isTest = 1
	}
	has, notice := c.NoticeService.GetNotice(isTest)
	notices := []models.Notice{}
	if has {
		notices = append(notices, *notice)
	}
	c.success(ctx, notices)
}
