package auth

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`        // 用户名
	Email    string `form:"email" binding:"required,email,min=3,max=100"` // 邮箱
	Password string `form:"password" binding:"required,min=8"`            // 密码
}
