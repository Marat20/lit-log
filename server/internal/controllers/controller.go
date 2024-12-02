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

	routes := r.Group("/books")
	routes.GET("/:id", h.GetBook)
	routes.GET("/", h.GetAllBooks)
	routes.POST("/", h.AddBook)
	routes.DELETE("/:id", h.DeleteBook)

	routes.GET("/progress/:id", h.GetReadingProgress)
	// routes.PATCH("/progress/:id", h.UpdateReadingProgress)
	routes.PATCH("/goal/:id", h.UpdateDailyGoal)
	// routes.PUT("/:id", h.UpdateArticle)
}
