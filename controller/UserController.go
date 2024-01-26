package controller

import (
	"journey-user/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImplementation struct {
	service *services.UserServiceImplementation
}

func NewUserController(s *services.UserServiceImplementation) *UserControllerImplementation {
	return &UserControllerImplementation{service: s}
}

func (user UserControllerImplementation) Get(c *gin.Context) {
	response := user.service.Get(c)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully get user", "users": response})
}

func (user UserControllerImplementation) Registration(c *gin.Context) {

	response, err := user.service.Registration(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": response})
}
