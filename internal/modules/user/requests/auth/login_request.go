package auth

// LoginRequest 登录请求结构
type LoginRequest struct {
	Email    string `form:"email" binding:"required,email,min=3,max=100"` // 邮箱
	Password string `form:"password" binding:"required,min=8"`            // 密码
}
