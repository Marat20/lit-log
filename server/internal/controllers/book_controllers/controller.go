package book_controllers

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

	r.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	routes := r.Group("/:userId")
	routes.GET("/init", h.initUser)
	routes.GET("/:bookId", h.getBook)
	routes.GET("/", h.getAllBooks)
	routes.POST("/", h.addBook)
	routes.DELETE("/:bookId", h.deleteBook)
	routes.PUT("/:bookId", h.updateCurrentPage)

}
