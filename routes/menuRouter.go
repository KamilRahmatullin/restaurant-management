package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yourusername/yourproject/controllers"
)

func MenuRouter(router *gin.Engine) {
	menuGroup := router.Group("/menus")
	{
		menuGroup.GET("/", controller.GetMenu())
		menuGroup.GET("/:menu_id", controller.GetMenus())
		menuGroup.POST("/", controller.CreateMenu())
		menuGroup.PATCH("/:menu_id", controller.UpdateMenu())
	}
}
