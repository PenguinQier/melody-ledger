package responses

import (
	"fmt"
	userModel "github.com/PenguinQier/melody-ledger/internal/modules/user/models"
)

// User 用户响应结构
type User struct {
	ID    uint   // 用户ID
	Image string // 用户头像
	Name  string // 用户名
	Email string // 邮箱
}

// Users 用户列表响应结构
type Users struct {
	Data []User
}

// ToUser 将用户模型转换为响应结构
func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name), // 生成用户头像
	}
}
