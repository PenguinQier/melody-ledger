package bootstrap

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/routes"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/config"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/database"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/html"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/routing"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/sessions"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	router := routing.GetRouter()

	sessions.Start(router)
	static.LoadStatic(router)

	router.Static("/public", "./public")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	html.LoadHTML(router)
	routes.RegisterRoutes(router)

	// router.Run(":8080")
	routing.Serve()
}
