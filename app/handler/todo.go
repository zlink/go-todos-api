package handler

import (
	"github.com/gin-gonic/gin"
	"api/database"
)

func All(ctx *gin.Context) {
	var todo interface{}
	database.DB.Table("todos").First(&todo)
}

func One(ctx *gin.Context) {

}

func Create(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}