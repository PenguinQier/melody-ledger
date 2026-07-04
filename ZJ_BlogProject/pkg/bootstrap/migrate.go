package bootstrap

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/database/migration"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/config"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
