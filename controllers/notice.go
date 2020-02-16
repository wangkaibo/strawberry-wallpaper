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
	statusParam := ctx.Query("is_publish")
	isPublish := 1
	if statusParam == "0" {
		isPublish = 0
	}
	notices, err := c.NoticeService.GetNotices(isPublish)
	fmt.Println(err)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}

func (c *NoticeController) NoticeList(ctx *gin.Context) {
	notices, err := c.NoticeService.GetNoticeList()
	fmt.Println(err)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}
