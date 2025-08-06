package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/database"
	"github.com/kamilrahmatullin/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection = database.OpenCollection(database.Client, "order")
var tableCollection = database.OpenCollection(database.Client, "table")

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
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		table := models.Table{}
		order := models.Order{}

		if err := ctx.BindJSON(&order); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(table); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(c, bson.M{"table_id": order.Table_id}).Decode(&table)
			if err != nil {
				msg := fmt.Sprintf("table was not found")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(c, order)
		if insertErr != nil {
			msg := fmt.Sprintf("order was not created")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		order := models.Order{}
		table := models.Table{}

		var updateObj primitive.D

		orderId := ctx.Param("order_id")

		if err := ctx.BindJSON(&order); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			if err := tableCollection.FindOne(c, bson.M{"table_id": order.Table_id}).Decode(&table); err != nil {
				msg := fmt.Sprintf("Table was not found")
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			updateObj = append(updateObj, bson.E{Key: "table_id", Value: order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.Updated_at})

		upsert := true
		filter := bson.M{"order_id": orderId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			c,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
