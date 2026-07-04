package helpers

import (
	UserRepository "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/repositories"
	UserResponse "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/responses"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User

	// 从 session 获取用户 ID
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	// 查询用户信息
	var userRepo = UserRepository.New()
	user := userRepo.FindByID(userID)

	// 如果用户不存在，返回空响应
	if user.ID == 0 {
		return response
	}

	return UserResponse.ToUser(user)
}
