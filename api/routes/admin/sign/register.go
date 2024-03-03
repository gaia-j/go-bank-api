package adminRoutes

import (
	"fmt"
	"github.com/gaia-j/go-bank-api/internal/structs"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var registerAdminData structs.RegisterAdmin

	if err := c.ShouldBindJSON(&registerAdminData); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := registerAdminData.User
	password := registerAdminData.Password
	cpf := registerAdminData.Cpf

	fmt.Println(user, password, cpf)

	//_, err = database.Db.Exec("insert into users(username,password) values ($1,$2)", body.Username, body.Password)
	//if err != nil {
	//	fmt.Println(err)
	//	ctx.AbortWithStatusJSON(400, "Couldn't create the new user.")
	//} else {
	//	ctx.JSON(http.StatusOK, "User is successfully created.")
	//}

	c.JSON(200, gin.H{
		"message": "venha cadastrar paizao",
	})

}
