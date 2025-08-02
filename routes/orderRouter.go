package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kamilrahmatullin/restaurant-management/controllers"
)

func OrderRoutes(router *gin.Engine) {
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", controller.GetOrders())
		orderGroup.GET("/:order_id", controller.GetOrder())
		orderGroup.POST("/", controller.CreateOrder())
		orderGroup.PATCH("/:order_id", controller.UpdateOrder())
	}
}
