package routes

import (
	controllers "restaurant-management-go/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:id", controllers.GetFood())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:id", controllers.UpdateFood())
}
