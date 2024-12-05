package controllers

import (
	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

type handler struct {
	DB *bolt.DB
}

func RegisterRoutes(r *gin.Engine, db *bolt.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books")
	routes.GET("/:id", h.getBook)
	routes.GET("/", h.getAllBooks)
	routes.POST("/", h.addBook)
	routes.DELETE("/:id", h.deleteBook)
	routes.PUT("/:id", h.updateCurrentPage)

}
