package routes

import (
	homeCtrl "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/home/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)
}
