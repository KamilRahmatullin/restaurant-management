package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kamilrahmatullin/restaurant-management/controllers"
)

func TableRoutes(incomingRouter *gin.Engine) {
	tableGroup := incomingRouter.Group("/tables")
	{
		tableGroup.GET("/", controller.GetTables())
		tableGroup.GET("/:table_id", controller.GetTable())
		tableGroup.POST("/", controller.CreateTable())
		tableGroup.PATCH("/:table_id", controller.UpdateTable())
	}
}
