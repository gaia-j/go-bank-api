package api

import (
	adminRoutes "github.com/gaia-j/go-bank-api/api/routes/admin/sign"
	userRoutes "github.com/gaia-j/go-bank-api/api/routes/user/sign"
	"github.com/gin-gonic/gin"
)

func Start() error {
	router := gin.Default()

	adminRoute := router.Group("/admin")
	{
		adminRoute.POST("/register", adminRoutes.Register)
		adminRoute.POST("/login", adminRoutes.Login)
	}

	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", userRoutes.RegisterUser)
		userRoute.POST("/login", userRoutes.LoginUser)
	}

	err := router.Run(":3000")
	return err
}
