package main

import (
	"strawberry-wallpaper/routers"
	"strawberry-wallpaper/db"
)

func main() {
	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	db.InitMySql()
	r.Run(":8080")
}