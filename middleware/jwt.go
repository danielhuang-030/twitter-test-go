package middleware

import (
	"app/model"

	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID uint   `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// @from https://github.com/eddycjy/go-gin-example/blob/master/middleware/jwt/jwt.go
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		splitToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		token := splitToken[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": "You do not have enough permissions to perform this operation.",
			})

			c.Abort()
			return
		}

		claims, err := parseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": "The certificate has expired.",
				})

				c.Abort()
				return
			default:
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": "Credential verification failed.",
				})

				c.Abort()
				return
			}
		}

		user, _ := model.FindUser(claims.UserID)
		c.Set("currentUser", user)

		c.Next()
	}
}

// @from https://github.com/eddycjy/go-gin-example/blob/master/pkg/util/jwt.go
func parseToken(token string) (*Claims, error) {
	var jwtSecret []byte
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
