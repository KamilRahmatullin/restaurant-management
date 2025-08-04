package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamilrahmatullin/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// creating a context with possibility to time out
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		// fetching food id from context
		menuId := ctx.Param("menu_id")

		menu := models.Menu{}

		// looking for menu id in database
		if err := menuCollection.FindOne(c, bson.M{"menu_id": menuId}).Decode(&menu); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu"})
			return
		}

		ctx.JSON(http.StatusOK, menu)
	}
}

func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// creating context with timeout
		c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// looking for all the menus in database
		result, err := menuCollection.Find(context.TODO(), bson.M{})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the menu items"})
			return
		}

		var allMenus []bson.M
		if err = result.All(c, &allMenus); err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, allMenus)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		menu := models.Menu{}

		if err := ctx.BindJSON(&menu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(menu); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		menu.ID = primitive.NewObjectID()
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Menu_id = menu.ID.Hex()

		result, insertError := menuCollection.InsertOne(c, menu)

		if insertError != nil {
			msg := fmt.Sprintf("food item was not created")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		menu := models.Menu{}
		menuId := ctx.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		if err := ctx.BindJSON(&menu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(menu.Start_Date, menu.End_Date, time.Now()) {
				msg := "kindly retype the time"
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_Date})

			if menu.Name != "" {
				updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
			}

			if menu.Category != "" {
				updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Category})
			}

			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_at})

			upsert := true

			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result, err := menuCollection.UpdateOne(
				c,
				filter,
				bson.D{
					bson.E{Key: "$set", Value: updateObj},
				},
				&opt,
			)

			if err != nil {
				msg := "Menu update failed"
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			ctx.JSON(http.StatusOK, result)
		}
	}
}

func inTimeSpan(startDate, endDate *time.Time, now time.Time) bool {
	return true
}
