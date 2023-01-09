package middlewares

import (
	models "Basic/Auth-Api/Models"
	token "Basic/Auth-Api/Token"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleBaseAuthmiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized 1")
			c.Abort()
			return
		}
		uid, errs := token.ExtractTokenID(c)

		user, err := models.GetUserByID(uid)
		log.Println(&user)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized 2")
			c.Abort()
			return
		}

		if errs != nil {
			c.String(http.StatusUnauthorized, "Unauthorized 2")
			c.Abort()
			return
		}
		if strings.ToLower(user.Role.Name) != strings.ToLower(role) {
			c.String(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}
