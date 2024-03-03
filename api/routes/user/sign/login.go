package userRoutes

import (
	"context"
	"fmt"
	"github.com/gaia-j/go-bank-api/internal/database"
	"github.com/gaia-j/go-bank-api/internal/structs"
	"github.com/gaia-j/go-bank-api/internal/token"
	"github.com/gaia-j/go-bank-api/internal/validation/signValidation"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {

	var loginData structs.LoginUser
	var userData structs.LoginUser
	var userId int

	if err := c.ShouldBindJSON(&loginData); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := signValidation.ValidateLoginUserData(&loginData); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := database.Db.QueryRow(
		context.Background(),
		"SELECT id,email,password FROM users WHERE email = $1", loginData.Email).Scan(&userId, &userData.Email, &userData.Password)

	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{
			"error": "email ou senha incorretos",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginData.Password)); err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{
			"error": "email ou senha incorretos",
		})
		return
	}

	userToken := token.GenerateUserToken(userId, userData.Email)

	errToken := token.ValidateUserToken(userToken)

	if errToken != nil {
		fmt.Println(errToken)
		c.JSON(400, gin.H{
			"error": errToken.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": userToken,
	})

}
