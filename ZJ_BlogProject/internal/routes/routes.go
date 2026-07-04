package routes

import (
	article "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/article/routes"
	home "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/home/routes"
	user "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	home.Routes(router)
	user.Routes(router)
	article.Routes(router)
}
