package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/env"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	port := env.GetValue("PORT", "8000")

	router := gin.new()
	router.Use(gin.Logger())
}
