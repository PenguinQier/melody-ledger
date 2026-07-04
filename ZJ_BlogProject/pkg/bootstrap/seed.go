package bootstrap

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/database/seeder"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/config"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
