package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kamilrahmatullin/restaurant-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	foodGroup := incomingRoutes.Group("/foods")
	{
		foodGroup.GET("/", controller.GetFoods())
		foodGroup.GET("/:food_id", controller.GetFood())
		foodGroup.POST("", controller.CreateFood())
		foodGroup.PATCH("/:food_id", controller.UpdateFood())
	}
}
