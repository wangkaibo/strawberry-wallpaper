package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type BaseController struct {
	DB *sql.DB
}

func (b BaseController) success(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"ec": 0,
		"em": "success",
		"data": data,
	})
}
