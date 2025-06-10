package routes

import (
	"udbytter/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPositionRoutes(rg *gin.RouterGroup) {
	position := rg.Group("/position")
	{
		position.GET("", controllers.GetPositions)
		position.POST("", controllers.CreatePosition)
		position.DELETE("/:id", controllers.DeletePosition)
		position.PUT("/:id", controllers.UpdatePosition)
		position.GET("/average-yield", controllers.GetAverageYield)
		position.GET("/monthly-income", controllers.GetMonthlyIncome)
	}
}
