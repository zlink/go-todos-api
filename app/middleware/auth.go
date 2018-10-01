package middleware

import (
	"api/app/service/e"
	"api/app/service/config"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserName string `json:"username`
	Password string `json:"password"`
	jwt.StandardClaims
}

var secret = []byte(config.App.Secret)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			tokenClaims, err := jwt.ParseWithClaims(token, &Cliaims{}, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})

			if tokenClaims != nil {
				if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
					return claims, nil
				}
			}
		}

		ctx.Next()
	}

	return nil, err
}
