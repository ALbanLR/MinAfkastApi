package routes

import (
	"udbytter/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaxRateRoutes(rg *gin.RouterGroup) {
	taxRate := rg.Group("/taxRate")
	{
		taxRate.GET("", controllers.GetTaxRate)
	}
}
