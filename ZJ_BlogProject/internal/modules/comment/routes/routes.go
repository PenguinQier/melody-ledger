package routes

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/middlewares"
	commentCtrl "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/comment/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	commentsController := commentCtrl.New()

	authGroup := router.Group("/comments")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("", commentsController.Store)
	}
}
