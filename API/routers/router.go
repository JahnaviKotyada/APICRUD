package routers

import (
	"crudapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	return r
}
