package main

import (
	"strawberry-wallpaper/routers"
	"strawberry-wallpaper/db"
	"strawberry-wallpaper/bootstrap"
)

func main() {
	app := bootstrap.NewApp("strawberry")
	app.Configure(routers.SetupRouter)
	// Listen and Server in 0.0.0.0:8080
	db.InitMySql()
	app.Engine.Run(":8080")
}