package userRoutes

import (
	"fmt"
	"github.com/gaia-j/go-bank-api/internal/database"
	"github.com/gaia-j/go-bank-api/internal/structs"
	"github.com/gaia-j/go-bank-api/internal/validation/signValidation"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {

	var registerUserData structs.RegisterUser

	var createdUserId int

	if err := c.ShouldBindJSON(&registerUserData); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := signValidation.ValidateRegisterUserData(&registerUserData); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUserId, err := database.RegisterUser(registerUserData)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.CreateUserAccount(createdUserId); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.CreateAddress(createdUserId, registerUserData); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Usuário registrado com sucesso",
	})

}
