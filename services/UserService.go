package services

import (
	"journey-user/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db}
}

func (userService *UserServiceImpl) Get(c *gin.Context) []model.User {
	var users []model.User

	userService.db.Find(&users)

	return users
}
