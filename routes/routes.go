package routes

import (
	"api/app/handler"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.New()
	r.GET("/auth", handler.Login)

	return r
}
