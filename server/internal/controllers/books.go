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
	TotalPages uint   `json:"totalPages" binding:"required"`
	DailyGoal  uint   `json:"dailyGoal" binding:"required"`
}

type updateBookInput struct {
	PagesRead uint `json:"pagesRead" binding:"required"`
}

func (h handler) getBook(context *gin.Context) {
	id := context.Param("id")

	var book books.Book
	if err := h.DB.Where("ID = ?", id).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"book": book})
}

func (h handler) getAllBooks(context *gin.Context) {
	var books []books.Book
	result := h.DB.Find(&books)
	if result.Error != nil {
		context.AbortWithError(http.StatusNotFound, result.Error)
	}

	context.JSON(http.StatusOK, gin.H{"books": books})
}

func (h handler) addBook(context *gin.Context) {
	var input createBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := gonanoid.New()

	book := books.Book{
		ID:          id,
		Title:       input.Title,
		Author:      input.Author,
		TotalPages:  input.TotalPages,
		DailyGoal:   input.DailyGoal,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CurrentPage: 0,
		IsActive:    true,
		IsDone:      false,
		PagesRead:   0,
	}

	if err := h.DB.Create(&book).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"book": book})
}

func (h handler) deleteBook(context *gin.Context) {
	bookID := context.Param("id")

	var book books.Book
	if err := h.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record does not exist"})
		return
	}

	if err := h.DB.Delete(&book).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (h handler) updateCurrentPage(context *gin.Context) {
	id := context.Param("id")

	var book books.Book
	if result := h.DB.Where("id = ?", id).First(&book); result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record does not exist"})
		return
	}

	var input updateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pagesRead := book.PagesRead + input.PagesRead
	isDailyGoalDone := pagesRead >= book.DailyGoal

	updates := map[string]interface{}{
		"pages_read": pagesRead,
		"updated_at": time.Now(),
	}

	if book.CurrentPage+input.PagesRead >= book.TotalPages {
		updates["is_active"] = false
		updates["is_done"] = true
		updates["finished_at"] = time.Now()
		updates["current_page"] = book.TotalPages
	} else {
		updates["current_page"] = book.CurrentPage + input.PagesRead
	}

	if err := h.DB.Model(&book).Updates(updates).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"isDailyGoalDone": isDailyGoalDone, "book": book})
}
