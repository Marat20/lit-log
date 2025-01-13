package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.Static("/assets", "./dist/assets")

	return r
}
