package bootstrap

import (
	"github.com/PenguinQier/melody-ledger/internal/database/migration"
	"github.com/PenguinQier/melody-ledger/pkg/config"
	"github.com/PenguinQier/melody-ledger/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
