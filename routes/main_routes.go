package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(app *gin.Engine) {
	api := app.Group("/api")
	RegisterPositionRoutes(api)
	RegisterTaxRateRoutes(api)
}
