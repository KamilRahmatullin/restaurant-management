package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
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
	return func(ctx *gin.Context) {}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
