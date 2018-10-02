package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register application routess
func Register() *gin.Engine {
	r := gin.New()

	r.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "OK"})
	})

	return r
}
