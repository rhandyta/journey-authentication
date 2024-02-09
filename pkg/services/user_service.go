package services

import (
	"errors"
	"journey-user/helper"
	models "journey-user/pkg/models"
	"journey-user/pkg/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserById(userId uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) CreateUser(user *models.User) error {

	existingUser, err := us.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with the same email already exists")
	}

	password, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	return us.userRepo.CreateUser(user)
}

func (us *userService) GetUserById(userId uint) (*models.User, error) {
	return us.userRepo.GetUserById(userId)
}

func (us *userService) GetUserByEmail(email string) (*models.User, error) {
	return us.userRepo.GetUserByEmail(email)
}
