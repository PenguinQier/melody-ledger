package models

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/models"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:191"`
	Content string `gorm:"text"`
	Files   string `gorm:"text"`
	Image   string
	UserID  uint
	User    models.User
}
