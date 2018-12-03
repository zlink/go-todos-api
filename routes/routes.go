package routes

import (
	"api/app/service/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type YmCrmLeadsRecord struct {
	Id      uint `gorm:id`
	LeadsId uint `gorm:leads_id`
	OwnersUserId uint `gorm:owner_`
}

// Register application routess
func Register() *gin.Engine {
	r := gin.New()

	r.GET("/user", func(ctx *gin.Context) {
		database.DB.Rows()
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "OK"})
	})

	return r
}
