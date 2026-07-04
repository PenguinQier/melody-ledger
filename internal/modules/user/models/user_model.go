package models

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Name     string `gorm:"varchar:191"` // 用户名
	Email    string `gorm:"varchar:191"` // 邮箱
	Password string `gorm:"varchar:191"` // 密码(加密后的)
}
