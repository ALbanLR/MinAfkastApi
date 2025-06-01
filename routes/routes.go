package routes

import (
	"udbytter/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	api := app.Group("/api")
	{
		position := api.Group("/position")
		{
			position.GET("", controllers.GetPositions)
			position.POST("", controllers.CreatePosition)
			position.DELETE("/:id", controllers.DeletePosition)
			position.PUT("/:id", controllers.UpdatePosition)

			// Routes métiers spécifiques
			position.GET("/average-yield", controllers.GetAverageYield)
			position.GET("/monthly-income", controllers.GetMonthlyIncome)
		}
	}
}
