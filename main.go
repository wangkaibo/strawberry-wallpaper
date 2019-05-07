package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strawberry-wallpaper/bootstrap"
	"strawberry-wallpaper/routers"
)
func init() {
	// 加载环境变量
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	err := godotenv.Load( path + "/.env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	gin.SetMode(os.Getenv("RUN_MODE"))
	app := bootstrap.NewApp("strawberry")
	app.Configure(routers.SetupRouter)
	// Listen and Server in 0.0.0.0:8080
	err := app.Engine.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}