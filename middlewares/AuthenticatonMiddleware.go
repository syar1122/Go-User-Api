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
			log.Panic(err.Error())
			c.String(http.StatusUnauthorized, "Unauthorized 0")
			c.Abort()
			return
		}
		c.Next()
	}
}
