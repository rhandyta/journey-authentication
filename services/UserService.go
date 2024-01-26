package services

import (
	"journey-user/helper"
	"journey-user/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServiceImplementation struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImplementation {
	return &UserServiceImplementation{db: db}
}

func (userService *UserServiceImplementation) Get(c *gin.Context) []model.User {
	var users []model.User

	userService.db.Find(&users)

	return users
}

func (userService *UserServiceImplementation) Registration(c *gin.Context) (model.User, error) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		helper.NewAppError("Invalid Request Body", http.StatusBadRequest)
		c.Set("Error", err)
		return model.User{}, err
	}

	password, err := helper.HashPassword(newUser.Password)
	if err != nil {
		helper.NewAppError("Invalid Format Password", http.StatusBadRequest)
		c.Set("Error", err)
		return model.User{}, err
	}

	newUser.Password = password
	if err := userService.db.Create(&newUser).Error; err != nil {
		helper.NewAppError("User Registration Failed", http.StatusInternalServerError)
		c.Set("Error", err)
		return model.User{}, err
	}

	return newUser, nil
}
