package main

import (
	"strawberry/routers"
	"strawberry/db"
)

func main() {
	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	db.InitMySql()
	r.Run(":8080")
}