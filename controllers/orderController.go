package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/database"
	"github.com/kamilrahmatullin/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
)

var orderCollection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
			return
		}

		var allOrders []bson.M
		if err = result.All(c, &allOrders); err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, allOrders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		order := models.Order{}
		orderId := ctx.Param("order_id")
		filter := bson.M{"order_id": orderId}

		if err := orderCollection.FindOne(c, filter).Decode(&order); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch the order"})
			return
		}

		ctx.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
