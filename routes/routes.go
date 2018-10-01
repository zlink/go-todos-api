package routes

import (
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.New()

	r.GET("/user", func(ctx *gin.Context) {

	})

	return r
}
