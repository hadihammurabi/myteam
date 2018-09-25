package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/" || path == "/auth/signin" {
			c.Next()
		} else {
			if _, ok := c.Request.Header["Authorization"]; ok {
				auth := strings.Split(c.Request.Header["Authorization"][0], " ")
				if auth[0] != "Bearer" {
					c.JSON(400, gin.H{
						"message": "Unauthorized",
					})
					c.Abort()
				} else {
					thetoken := auth[1]
					token, _ := jwt.Parse(thetoken, func(t *jwt.Token) (interface{}, error) {
						return []byte("rahasiadong"), nil
					})
					if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
						c.Next()
					} else {
						c.JSON(400, gin.H{
							"message": "Unauthorized",
						})
						c.Abort()
					}
				}
			} else {
				c.JSON(400, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
			}
		}
	}
}