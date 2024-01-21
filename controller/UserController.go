package controller

import (
	"journey-user/helper"
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

	webResponse := helper.ReturnFormatJson{
		Message: "Successfully get user",
		Data:    response,
	}

	c.JSON(http.StatusOK, webResponse)
}
