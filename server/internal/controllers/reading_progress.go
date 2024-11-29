package controllers

import (
	"lit-log/internal/models/books"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type updateReadingProgressInput struct {
	Content string `json:"content" binding:"required"`
}

func (h handler) GetReadingProgress(context *gin.Context) {
	id := context.Param("id")

	var book books.Book
	if err := h.DB.Where("id = ?", id).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	h.DB.Find(&book)

	context.JSON(http.StatusOK, gin.H{"book": book})
}


func (h handler) UpdateReadingProgress(context *gin.Context) {
	var input createBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := gonanoid.New()
	book := books.Book{Id: id, Title: input.Title, Author: input.Author, CreatedAt: time.Now()}
	h.DB.Create(&book)

	context.JSON(http.StatusOK, gin.H{"book": book})
}

func (h handler) UpdateDailyGoal(context *gin.Context) {
	var input createBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := gonanoid.New()
	book := books.Book{Id: id, Title: input.Title, Author: input.Author, CreatedAt: time.Now()}
	h.DB.Create(&book)

	context.JSON(http.StatusOK, gin.H{"book": book})
}
