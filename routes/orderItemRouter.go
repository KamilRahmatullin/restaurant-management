package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	orderItemGroup := incomingRoutes.Group("/orderItems")
	{
		orderItemGroup.GET("/", controllers.GetOrderItems())
		orderItemGroup.GET("/:orderItem_id", controllers.GetOrderItem())
		orderItemGroup.POST("/", controllers.CreateOrderItem())
		orderItemGroup.PATCH("/:orderItem_id", controllers.UpdateOrderItem())
	}
}
