package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routesBook := r.Group("/books")
	routesBook.GET("/:id", h.GetBook)
	routesBook.GET("/", h.GetAllBooks)
	routesBook.POST("/", h.AddBook)
	routesBook.DELETE("/:id", h.DeleteBook)

	routesReadingProgress := r.Group("/books/progress")
	routesReadingProgress.GET("/:id", h.GetReadingProgress)
	routesReadingProgress.PATCH("/", h.UpdateReadingProgress)
	routesReadingProgress.PATCH("/:id", h.UpdateDailyGoal)
	// routes.PUT("/:id", h.UpdateArticle)
}
