package controllers

import (
	"LostPaws/backend/config"
	"LostPaws/backend/models"
	"net/http"

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
func GetDog(c *gin.Context) {
	id := c.Param("id")
	var dog models.Dog
	if err := config.DB.First(&dog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontro la mascota"})
		return
	}
	c.JSON(http.StatusOK, dog)
}
func UpdateDog(c *gin.Context) {
	id := c.Param("id")
	var dog models.Dog

	if err := config.DB.First(&dog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mascota no Encontrada"})
		return
	}
	var input struct {
		Estado bool `json:"estado"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dog.Estado = input.Estado
	config.DB.Save(&dog)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Estado actualizado", "dog": dog})
}
