package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func GenerateUserToken(userId int, userEmail string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	JwtPass := os.Getenv("JWT_PASS")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  userId,
		"email":   userEmail,
		"expires": time.Now().Add(time.Hour * 72).Unix(),
		"admin":   false,
	})

	if err != nil {
		fmt.Println("Erro ao gerar a chave secreta:", err)
	}

	userToken, err := token.SignedString([]byte(JwtPass))

	if err != nil {
		fmt.Println("Error generating token")
		panic(err)
	}

	return userToken

}
