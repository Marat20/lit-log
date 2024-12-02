package controllers

import (
	"lit-log/internal/models/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateDailyGoalInput struct {
	DailyGoal uint `json:"dailyGoal" binding:"required"`
}

func (h handler) GetReadingProgress(context *gin.Context) {
	id := context.Param("id")

	var progress books.ReadingProgress
	if err := h.DB.Where("id = ?", id).First(&progress).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	h.DB.Find(&progress)

	context.JSON(http.StatusOK, gin.H{"progress": progress})
}

// func (h handler) UpdateReadingProgress(context *gin.Context) {
// 	id := context.Param("id")

// 	var progress books.ReadingProgress
// 	if err := h.DB.Where("id = ?", id).First(&progress).Error; err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
// 		return
// 	}

// 	var input updateDailyGoalInput
// 	if err := context.ShouldBindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	h.DB.Model(&article).Update("content", input)
// 	context.JSON(http.StatusOK, gin.H{"article": article})
// }

func (h handler) UpdateDailyGoal(context *gin.Context) {
	id := context.Param("id")

	var progress books.ReadingProgress
	if err := h.DB.Where("id = ?", id).First(&progress).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input updateDailyGoalInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&progress).Update("daily_goal", input)
	context.JSON(http.StatusOK, gin.H{"progress": progress})
}
