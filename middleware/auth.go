package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strawberry-wallpaper/controllers"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("TOKEN")
		tokenList := strings.Split(token, ".")
		fmt.Println(len(tokenList))
		if len(tokenList) != 3 {
			noAuth(ctx)
		}
		header := tokenList[0]
		payload := tokenList[1]
		reqSign := tokenList[2]
		sign := controllers.GetJwtSign(header, payload)
		if reqSign != sign {
			noAuth(ctx)
		}
		ctx.Next()
	}
}

func noAuth(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 401,
		"message": "无权限",
		"data": gin.H{},
	})
}