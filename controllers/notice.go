package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	notices, err := c.NoticeService.GetNotices(isTest)
	fmt.Println(err)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}
