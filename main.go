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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Ganti sesuai frontend
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Jika OPTIONS, langsung return
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Todo{})
	router := gin.Default()

	router.Use(loggerMiddleware())
	router.Use(CORSMiddleware())
	router.GET("/", controllers.Home)
	router.GET("/lists", controllers.GetLists)
	router.POST("/lists", controllers.CreateList)
	router.DELETE("lists/:id", controllers.DeleteList)
	router.PUT("lists/:id", controllers.UpdateList)
	router.Run("localhost:8000")
}
