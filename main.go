package main

import (
	"strawberry-wallpaper/bootstrap"
	"strawberry-wallpaper/routers"
)

func main() {
	app := bootstrap.NewApp("strawberry")
	app.Configure(routers.SetupRouter)
	// Listen and Server in 0.0.0.0:8080
	app.Engine.Run(":8080")
}