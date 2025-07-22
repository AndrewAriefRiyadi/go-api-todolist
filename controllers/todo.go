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

func CreateList(ctx *gin.Context) {

	var todo models.Todo

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    todo,
	})
}

func DeleteList(ctx *gin.Context) {
	id := ctx.Param("id")

	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Todo berhasil di-delete",
	})
}

func UpdateList(ctx *gin.Context) {
	id := ctx.Param("id")

	var todo models.Todo
	var todoBody models.Todo

	// Find based on ID
	if err := config.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Initialize todo based on json body
	if err := ctx.ShouldBindJSON(&todoBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Title = todoBody.Title
	todo.Completed = todoBody.Completed

	if err := config.DB.Save(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    todo,
	})
}
