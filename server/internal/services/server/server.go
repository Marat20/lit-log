package server

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	return r
}
