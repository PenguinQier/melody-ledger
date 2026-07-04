package services

import (
	"github.com/PenguinQier/melody-ledger/internal/modules/user/requests/auth"
	UserResponse "github.com/PenguinQier/melody-ledger/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	UpdateProfile(userID uint, request auth.ProfileRequest) (UserResponse.User, error)
}
