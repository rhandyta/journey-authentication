package controller

import (
	"journey-user/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserControllerImpl struct {
	service *services.UserServiceImpl
}

func NewUserController(db *gorm.DB) *UserControllerImpl {
	service := services.NewUserService(db)
	return &UserControllerImpl{service: service}
}

func (user UserControllerImpl) Get(c *gin.Context) {
	response := user.service.Get(c)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully get user", "users": response})
}
