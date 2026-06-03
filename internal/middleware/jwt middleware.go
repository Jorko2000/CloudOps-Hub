package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

func JWTAuth(
	secret string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader(
			"Authorization",
		)

		if authHeader == "" {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing token",
				},
			)

			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwtlib.Parse(
			tokenString,
			func(
				token *jwtlib.Token,
			)(interface{}, error) {

				return []byte(secret), nil
			},
		)

		if err != nil || !token.Valid {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			return
		}

		claims := token.Claims.(jwtlib.MapClaims)

		c.Set(
			"email",
			claims["email"],
		)

		c.Next()
	}
}
