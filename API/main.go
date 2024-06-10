package main

import (
	"crudapi/models"
	"crudapi/routers"
)

func main() {
	models.ConnectDatabase()
	r := routers.SetUpRouter()
	r.Run(":8080")
}
