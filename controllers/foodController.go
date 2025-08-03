package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kamilrahmatullin/restaurant-management/database"
	"github.com/kamilrahmatullin/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// creating a context with possibility to time out
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		// fetching food id from context
		foodId := ctx.Param("food_id")

		food := models.Food{}

		// looking for food id in database
		if err := foodCollection.FindOne(c, bson.M{"food_id": foodId}).Decode(&food); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
			return
		}

		ctx.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// creating a context with possibility to time out
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		menu := models.Menu{}
		food := models.Food{}

		// binding data from context to food model
		if err := ctx.BindJSON(&food); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if validationErr := validate.Struct(food); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// look for the food's menu
		if err := menuCollection.FindOne(c, bson.M{"menu_id": food.Menu_id}).Decode(&menu); err != nil {
			msg := fmt.Sprintf("menu was not found")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// creating a food object and inserting it into database
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		num := toFixed(*food.Price, 2)
		food.Price = &num

		result, insertError := foodCollection.InsertOne(ctx, food)

		if insertError != nil {
			msg := fmt.Sprintf("food item was not created")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}
