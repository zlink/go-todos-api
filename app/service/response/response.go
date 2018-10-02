package response

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

func (g *Gin) Success(code int, msg string, data interface{}) {
	g.Ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	return
}

func (g *Gin) Error(code int, msg string, data interface{}) {
	g.Ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	return
}
