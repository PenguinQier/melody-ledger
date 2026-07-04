package bootstrap

import (
	"github.com/PenguinQier/melody-ledger/internal/database/seeder"
	"github.com/PenguinQier/melody-ledger/pkg/config"
	"github.com/PenguinQier/melody-ledger/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
