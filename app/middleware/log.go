package middleware

import "github.com/gin-gonic/gin"

type Logger struct {
}

func (logger *Logger) Notice() gin.HandlerFunc {

}

func (logger *Logger) Info() gin.HandlerFunc {

}

func (logger *Logger) Warn() gin.HandlerFunc {

}

func (logger *Logger) Error() gin.HandlerFunc {

}
