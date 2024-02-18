package services

import (
	"fmt"
	"journey-user/helper"
	"journey-user/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationServiceImpl struct {
	db *gorm.DB
}

func NewAuthenticationService(db *gorm.DB) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{db: db}
}

func (userService *UserServiceImpl) Registration(user model.User) (model.User, error) {
	password, err := helper.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}

	user.Password = password
	if err := userService.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (userService *UserServiceImpl) Login(request model.Login) (model.AuthUser, error) {
	var authUser model.AuthUser
	var user model.User

	userService.db.Where("username = ?", request.Username).
		Or("email = ? ", request.Email).
		First(&user)

	if user.Username == "" {
		return authUser, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return authUser, err
	}

	tokenString, err := createToken(user)
	if err != nil {
		return authUser, nil
	}
	authUser.User = user
	authUser.Token = tokenString
	return authUser, nil
}

func createToken(user model.User) (string, error) {
	claims := model.AuthUserClaim{
		user.Username,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secred-key"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
