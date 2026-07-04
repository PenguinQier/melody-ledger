package migrations

import (
	"github.com/PenguinQier/melody-ledger/internal/modules/comment/models"

	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Comment{})
}

func CreateArticlesTable(db *gorm.DB) error {
	return db.Exec("ALTER TABLE articles ADD COLUMN IF NOT EXISTS files TEXT AFTER content").Error
}

func CreateCommentsTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Comment{})
}
