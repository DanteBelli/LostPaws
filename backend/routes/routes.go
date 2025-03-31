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
	}
	return r
}
