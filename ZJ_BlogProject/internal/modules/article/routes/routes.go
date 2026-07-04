package routes

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/middlewares"
	articleCtrl "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/controllers"
	commentCtrl "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/comment/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := articleCtrl.New()
	commentsController := commentCtrl.New()

	router.GET("/articles/:id", articlesController.Show)
	router.GET("/search", articlesController.Search)
	router.GET("/api/search", articlesController.APISearch)

	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", articlesController.Create)
		authGroup.POST("/store", articlesController.Store)
		authGroup.GET("/:id/edit", articlesController.Edit)
		authGroup.POST("/:id/update", articlesController.HandleUpdate)
		authGroup.POST("/:id/delete", articlesController.HandleDelete)
		authGroup.POST("/:id/cover", articlesController.UpdateCover)
	}

	commentGroup := router.Group("/comments")
	commentGroup.Use(middlewares.IsAuth())
	{
		commentGroup.POST("", commentsController.Store)
		commentGroup.POST("/:id/delete", commentsController.HandleDelete)
	}
}
