package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/env"
	"github.com/kamilrahmatullin/restaurant-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	port := env.GetValue("PORT", "8000")

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
}
