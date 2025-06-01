package controllers

import (
	"net/http"
	"udbytter/models"
	services "udbytter/service"

	"github.com/gin-gonic/gin"
)

func GetPositions(c *gin.Context) {
	positions, err := services.GetAllPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch positions"})
		return
	}
	c.JSON(http.StatusOK, positions)
}

func CreatePosition(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreatePosition(&position); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create position"})
		return
	}
	c.JSON(http.StatusCreated, position)
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeletePosition(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Position deleted"})
}

func UpdatePosition(c *gin.Context) {
	id := c.Param("id")
	var input models.Position
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := services.UpdatePosition(id, &input)
	if err != nil {
		if err.Error() == "position not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update position"})
		}
		return
	}

	c.JSON(http.StatusOK, updated)
}
