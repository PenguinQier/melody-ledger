package middlewares

import (
	UserRepository "github.com/PenguinQier/melody-ledger/internal/modules/user/repositories"
	"github.com/PenguinQier/melody-ledger/pkg/sessions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		// 请求处理前的认证检查

		c.Next()
	}
}
