package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "auth successed"})
}

func Logout(ctx *gin.Context) {

}

func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "pong"})
}
