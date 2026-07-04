package models

import (
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/models"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:text"`
	UserID    uint
	User      models.User
	ArticleID uint
}
