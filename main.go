package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/database"
	"github.com/kamilrahmatullin/restaurant-management/env"
	"github.com/kamilrahmatullin/restaurant-management/middleware"
	"github.com/kamilrahmatullin/restaurant-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := env.GetValue("PORT", "8000")

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
