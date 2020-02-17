package routers

import (
	"github.com/gin-contrib/cors"
	"strawberry-wallpaper/bootstrap"
	"strawberry-wallpaper/controllers"
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
	r.GET("/notice_list", noticeController.NoticeList)
}