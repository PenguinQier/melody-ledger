package services

import (
	"errors"
	userModels "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/models"
	UserRepository "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/repositories"
	"github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/requests/auth"
	UserResponse "github.com/PenguinQier/melody-ledger/ZJ_BlogProject/internal/modules/user/responses"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModels.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating the user")
	}

	return UserResponse.ToUser(user), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)

	if user.ID != 0 {
		return true
	}
	return false
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User
	existsUser := userService.userRepository.FindByEmail(request.Email)

	if existsUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return UserResponse.ToUser(existsUser), nil
}

func (userService *UserService) UpdateProfile(userID uint, request auth.ProfileRequest) (UserResponse.User, error) {
	var response UserResponse.User
	user := userService.userRepository.FindByID(int(userID))

	if user.ID == 0 {
		return response, errors.New("user not found")
	}

	// Check if email is already taken by another user
	if user.Email != request.Email {
		existingUser := userService.userRepository.FindByEmail(request.Email)
		if existingUser.ID != 0 && existingUser.ID != user.ID {
			return response, errors.New("email already exists")
		}
	}

	user.Name = request.Name
	user.Email = request.Email

	// Only update password if provided
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
		if err != nil {
			return response, errors.New("error hashing the password")
		}
		user.Password = string(hashedPassword)
	}

	updatedUser := userService.userRepository.Update(user)
	return UserResponse.ToUser(updatedUser), nil
}
