package main

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/database/migrations"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/database"
	"log"
)

func main() {
	db := database.Connection()

	err := migrations.Migrate(db)
	if err != nil {
		log.Fatal("Error running migrations:", err)
	}

	log.Println("Migrations completed successfully")
}
