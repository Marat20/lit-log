package controllers

import (
	"lit-log/internal/models/books"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type createBookInput struct {
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	TotalPages string `json:"totalPages" binding:"required"`
}

type updateBookInput struct {
	Content string `json:"content" binding:"required"`
}

func (h handler) GetBook(context *gin.Context) {
	id := context.Param("id")

	var book books.Book
	if err := h.DB.Where("id = ?", id).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	h.DB.Find(&book)

	context.JSON(http.StatusOK, gin.H{"book": book})
}

func (h handler) GetAllBooks(context *gin.Context) {
	var books []books.Book
	result := h.DB.Find(&books)
	if result.Error != nil {
		context.AbortWithError(http.StatusNotFound, result.Error)
	}

	context.JSON(http.StatusOK, gin.H{"books": books})
}

func (h handler) AddBook(context *gin.Context) {
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

func (h handler) DeleteBook(context *gin.Context) {
	id := context.Param("id")

	var book books.Book
	if err := h.DB.Where("id = ?", id).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	h.DB.Delete(&book)

	context.JSON(http.StatusOK, gin.H{"success": true})
}

// func (h handler) UpdateArticle(context *gin.Context) {
// 	id := context.Param("id")

// 	var article article.Article
// 	if err := h.DB.Where("id = ?", id).First(&article).Error; err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
// 		return
// 	}

// 	var input updateArticleInput
// 	if err := context.ShouldBindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	h.DB.Model(&article).Update("content", input)
// 	context.JSON(http.StatusOK, gin.H{"article": article})
// }
