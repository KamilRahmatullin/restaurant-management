package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/kamilrahmatullin/restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userGroup := incomingRoutes.Group("/users")
	{
		userGroup.GET("/", controller.GetUsers())
		userGroup.GET("/:user_id", controller.GetUser())
		userGroup.POST("/signup", controller.SignUp())
		userGroup.POST("/login", controller.Login())
	}
}
