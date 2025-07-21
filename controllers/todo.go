package controllers

import (
	"net/http"

	"github.com/AndrewAriefRiyadi/gin-be/config"
	"github.com/AndrewAriefRiyadi/gin-be/models"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to ToDo-Lists in GO")
}

func GetLists(ctx *gin.Context) {
	var todos []models.Todo
	err := config.DB.Find(&todos).Error

	//error handling get todos
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//Cek apakah list ada atau tidak
	if len(todos) == 0 {
		ctx.JSON(500, gin.H{"message": "Tidak ada Lists"})
		return
	}
	ctx.JSON(200, todos)
}
