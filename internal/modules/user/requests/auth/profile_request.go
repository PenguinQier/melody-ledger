package auth

type ProfileRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	Email    string `form:"email" binding:"required,email,min=3,max=100"`
	Password string `form:"password"`
}
