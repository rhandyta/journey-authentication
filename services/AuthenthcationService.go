package services

import (
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

	userSaved := userService.db.Where("email = ? ", request.Identifier).
		Or("username = ? ", request.Identifier).
		First(&user)

	if userSaved == nil {
		return authUser, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return authUser, err
	}

	tokenString, err := createToken(request.Identifier)
	if err != nil {
		return authUser, nil
	}
	authUser.User = &user
	authUser.Token = tokenString
	return authUser, nil
}

func createToken(Identifier string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"identifier": Identifier,
			"expired_at": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString("jwt-auth")
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
