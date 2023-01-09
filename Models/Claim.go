package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	now := time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) {
		return nil
	}
	return fmt.Errorf("Token is not Valid")
}
