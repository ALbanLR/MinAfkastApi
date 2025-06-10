package controllers

import (
	"net/http"
	services "udbytter/service"

	"github.com/gin-gonic/gin"
)

func GetTaxRate(c *gin.Context) {
	taxRate, err := services.GetAllTaxRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tax rates"})
		return
	}
	c.JSON(http.StatusOK, taxRate)
}
