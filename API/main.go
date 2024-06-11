package main

import (
	"api/models"
	"api/routers"
)

func main() {
	models.ConnectDatabase()
	r := routers.SetUpRouter()
	r.Run(":8080")
}
