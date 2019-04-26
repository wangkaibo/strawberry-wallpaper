package routers

import (
	"strawberry-wallpaper/controllers"
	"strawberry-wallpaper/bootstrap"
	"strawberry-wallpaper/services"
)

func SetupRouter(b *bootstrap.Bootstrap) {
	r := b.Engine
	statisticService := services.NewStatisticService()
	statisticController := &controllers.StatisticController{
		Service: statisticService,
	}
	r.GET("/register", statisticController.Register)
}