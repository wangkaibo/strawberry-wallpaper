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
	status := 0
	if paramTest == "1" {
		status = 1
	}
	notices, err := c.NoticeService.GetNotices(status)
	fmt.Println(err)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}
