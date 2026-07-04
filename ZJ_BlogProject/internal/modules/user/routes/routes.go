package routes

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/middlewares"
	userCtrl "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/profile", userController.Profile)
		authGroup.POST("/profile", userController.HandleUpdateProfile)
		authGroup.POST("/logout", userController.HandleLogout)
	}
}
