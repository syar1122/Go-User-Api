package middlewares

import (
	token "Basic/Auth-Api/Token"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			log.Println(err.Error())
			return
		}
		c.Next()
	}
}
