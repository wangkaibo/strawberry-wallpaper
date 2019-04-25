package routers

import (
	"strawberry-wallpaper/controllers"
	"strawberry-wallpaper/bootstrap"
)

func SetupRouter(b *bootstrap.Bootstrap) {
	r := b.Engine
	statistic := new(controllers.StatisticController)
	r.GET("/test", statistic.Statistic)
}