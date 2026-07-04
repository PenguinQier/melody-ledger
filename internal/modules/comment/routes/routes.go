package routes

import (
	"github.com/PenguinQier/melody-ledger/internal/middlewares"
	commentCtrl "github.com/PenguinQier/melody-ledger/internal/modules/comment/controllers"

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
