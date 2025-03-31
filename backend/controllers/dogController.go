package controllers

import (
	"net/http"

	"github.com/DanteBelli/LostPaws/config"
	"github.com/DanteBelli/LostPaws/models"
	"github.com/gin-gonic/gin"
)

func CreateDog(c *gin.Context) {
	var dog models.Dog
	if err := c.ShouldBindJSON(&dog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&dog)
	c.JSON(http.StatusOK, gin.H{"mensaje": " Mascota ingresada", "dog": dog})
}
func GetDogs(c *gin.Context) {
	var dogs []models.Dog
	config.DB.Find(&dogs)
	c.JSON(http.StatusOK, dogs)
}
