package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strawberry-wallpaper/controllers"
	"strings"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		tokenList := strings.Split(token, ".")
		if len(tokenList) != 3 {
			noAuth(ctx)
		}
		header := tokenList[0]
		payload := tokenList[1]
		reqSign := tokenList[2]
		sign := controllers.GetJwtSign(header, payload)
		payloadStr,err := base64.StdEncoding.DecodeString(payload)
		var payLoadMap map[string]interface{}
		err = json.Unmarshal(payloadStr, &payLoadMap)
		unixTime := float64(time.Now().Unix())
		if payLoadMap["exp"].(float64) < unixTime {
			noAuth(ctx)
		}
		if err != nil || reqSign != sign {
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