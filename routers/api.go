package routers

import (
	"github.com/gin-gonic/gin"
	"strawberry/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	statistic := new(controllers.StatisticController)
	r.GET("/test", statistic.Statistic)
	return r
}