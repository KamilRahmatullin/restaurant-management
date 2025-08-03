package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/database"
	"github.com/kamilrahmatullin/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		foodId := ctx.Param("food_id")

		food := models.Food{}

		if err := foodCollection.FindOne(c, bson.M{"food_id": foodId}).Decode(&food); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}

		ctx.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}
