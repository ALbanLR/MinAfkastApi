package routes

import (
	"udbytter/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/api/position", controllers.GetPositions)
	app.POST("/api/position", controllers.CreatePosition)
	app.DELETE("/api/position/:id", controllers.DeletePosition)
	app.PUT("/api/position/:id", controllers.UpdatePosition)
}
