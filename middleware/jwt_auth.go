package middleware

import (
	util "api/util/jwt"
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

			usersId, err := util.ParseToken(token)
			c.Set("userssId", usersId)
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
