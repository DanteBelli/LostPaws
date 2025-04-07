package routes

import (
	"LostPaws/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/dogs", controllers.CreateDog)
		api.GET("/dogs", controllers.GetDogs)
		api.GET("/dogs/:id", controllers.GetDog)
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.PUT("/dogs/:id", controllers.UpdateDog)
	}
	return r
}
