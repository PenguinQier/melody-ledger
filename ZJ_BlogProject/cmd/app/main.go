package main

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 注册所有路由
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
