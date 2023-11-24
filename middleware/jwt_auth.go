package middleware

import (
	"api/util/jwt_secret"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var isPass = true
		var message string

		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			message = "token not found"
			isPass = false
		} else {
			jwtToken := strings.Split(bearerToken, " ")
			token := jwtToken[1]

			claims, err := jwt_secret.ParseToken(token)
			if claims != nil{
				c.Set("usersId", claims.UserID)
			}
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					message = "token expired"
					isPass = false
				default:
					message = "token fail"
					isPass = false
				}
			}
		}

		if !isPass {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": message,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
