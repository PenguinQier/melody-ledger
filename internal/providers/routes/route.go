package routes

import (
	articleRoutes "github.com/PenguinQier/melody-ledger/internal/modules/article/routes"
	homeRoutes "github.com/PenguinQier/melody-ledger/internal/modules/home/routes"
	userRoutes "github.com/PenguinQier/melody-ledger/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
