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

	books := r.Group("/books")
	books.GET("/:id", h.getBook)
	books.GET("/", h.getAllBooks)
	books.POST("/", h.addBook)
	books.DELETE("/:id", h.deleteBook)
	books.PUT("/:id", h.updateCurrentPage)

}
