package adminRoutes

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "venha logar paizao",
	})

}
