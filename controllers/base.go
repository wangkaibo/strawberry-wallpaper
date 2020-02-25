package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type BaseController struct {}

func (b *BaseController) success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": data,
	})
}

func (b *BaseController) error(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"message": message,
		"data": data,
	})
}
