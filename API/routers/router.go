package routers

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.POST("/addresses", controllers.CreateAddress)
	r.GET("/addresses", controllers.GetAddresses)
	r.GET("/addresses/:id", controllers.GetAddressByID)
	r.PUT("/addresses/:id", controllers.UpdateAddress)
	r.DELETE("/addresses/:id", controllers.DeleteAddress)
	return r
}
