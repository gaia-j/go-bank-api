package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ValidateUserToken(userToken string) error {
	JwtPass := os.Getenv("JWT_PASS")
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtPass), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
