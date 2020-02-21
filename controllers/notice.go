package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strawberry-wallpaper/services"
	"strconv"
	"time"
)

type NoticeController struct {
	BaseController
	NoticeService services.NoticeService
}

func (c *NoticeController) Notice(ctx *gin.Context) {
	notices, err := c.NoticeService.GetNotices()
	fmt.Println(err)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}

func (c *NoticeController) NoticeList(ctx *gin.Context) {
	notices, err := c.NoticeService.GetNoticeList()
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}

func (c *NoticeController) DeleteNotice(ctx *gin.Context) {
	id := ctx.Param("id")
	notices, err := c.NoticeService.DeleteNotice(id)
	fmt.Println(err)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, notices)
	}
}


func (c *NoticeController) ChangeStatus(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	status := ctx.Request.FormValue("status")
	fmt.Println(status)
	if status == "" {
		c.error(ctx, 400, "参数错误", gin.H{})
		return
	}
	isPublish, _ := strconv.Atoi(status)
	err := c.NoticeService.ChangeStatus(id, isPublish)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, gin.H{})
	}
}

func (c *NoticeController) AddNotice(ctx *gin.Context) {
	content := ctx.Request.FormValue("content")
	if content == "" {
		c.error(ctx, 400, "内容不能为空", gin.H{})
		return
	}

	publishTimeStr := ctx.Request.FormValue("publish_time")
	expireTimeStr := ctx.Request.FormValue("expire_time")

	var err error
	publishTime, err := stringToTime(publishTimeStr)
	expireTime, err := stringToTime(expireTimeStr)
	if err != nil {
		c.error(ctx, 400, "参数错误", gin.H{})
		return
	}
	err = c.NoticeService.AddNotice(content, publishTime, expireTime)
	if err != nil {
		c.error(ctx, 500, err.Error(), gin.H{})
	} else {
		c.success(ctx, gin.H{})
	}
}

func stringToTime(ms string) (time.Time, error) {
	msInt,err := strconv.Atoi(ms)
	if err != nil {
		return time.Time{}, err
	}
	nsImt64 := int64(msInt) * int64(time.Millisecond)
	return time.Unix(0, nsImt64), nil
}
