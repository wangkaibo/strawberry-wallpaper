package routers

import (
	"github.com/gin-contrib/cors"
	"strawberry-wallpaper/bootstrap"
	"strawberry-wallpaper/controllers"
	"strawberry-wallpaper/middleware"
	"strawberry-wallpaper/services"
)

func SetupRouter(b *bootstrap.Bootstrap) {
	r := b.Engine
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	statisticService := services.NewStatisticService()
	statisticController := &controllers.StatisticController{
		StatisticService: statisticService,
	}
	noticeService := services.NewNoticeService()
	noticeController := &controllers.NoticeController{
		NoticeService: noticeService,
	}
	r.POST("/register", statisticController.Register)
	r.POST("/active", statisticController.Active)
	r.GET("/statistic", statisticController.Index)
	r.GET("/notice", noticeController.Notice)
	// 公告列表
	r.GET("/notice_list",middleware.Auth(), noticeController.NoticeList)
	// 删除公告 id放到 path
	r.DELETE("/notice/:id", middleware.Auth(), noticeController.DeleteNotice)
	// 更改状态 PATCH id放到 path，参数status：0-未发布，1-已发布
	r.PATCH("/notice/:id", middleware.Auth(), noticeController.ChangeStatus)
	// 增加公告 POST 参数：content：公告内容；publish_time：发布时间；expire_time：过期时间
	r.POST("/notice", middleware.Auth(), noticeController.AddNotice)
	r.POST("/login", noticeController.Login)
}