package migrations

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := CreateUsersTable(db); err != nil {
		return err
	}

	if err := CreateArticlesTable(db); err != nil {
		return err
	}

	if err := CreateCommentsTable(db); err != nil {
		return err
	}

	return nil
}
