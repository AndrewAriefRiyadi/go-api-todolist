package main

import (
	"log"
	"time"

	"github.com/AndrewAriefRiyadi/gin-be/config"
	"github.com/AndrewAriefRiyadi/gin-be/controllers"
	"github.com/AndrewAriefRiyadi/gin-be/models"
	"github.com/gin-gonic/gin"
)

func loggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration %v", ctx.Request.Method, ctx.Writer.Status(), duration)
	}
}

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Todo{})
	router := gin.Default()

	router.Use(loggerMiddleware())
	router.GET("/", controllers.Home)
	router.GET("/lists", controllers.GetLists)
	router.POST("/lists", controllers.CreateList)
	router.DELETE("lists/:id", controllers.DeleteList)
	router.PUT("lists/:id", controllers.UpdateList)
	router.Run("localhost:8000")
}
