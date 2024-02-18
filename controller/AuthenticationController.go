package controller

import (
	"journey-user/helper"
	"journey-user/model"
	"journey-user/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthenticationControllerImpl struct {
	service *services.UserServiceImpl
}

func NewAuthenticationController(db *gorm.DB) *AuthenticationControllerImpl {
	services := services.NewUserService(db)
	return &AuthenticationControllerImpl{service: services}
}

func (auth *AuthenticationControllerImpl) Registration(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	validationString := helper.RequestValidationHelper(user)

	if validationString != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationString})
		return
	}

	registUser, err := auth.service.Registration(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": registUser})
}

func (auth *AuthenticationControllerImpl) Login(c *gin.Context) {
	var userLogin model.Login
	if err := c.BindJSON(&userLogin); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	validationString := helper.RequestValidationHelper(userLogin)

	if validationString != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationString})
		return
	}

	authUser, err := auth.service.Login(userLogin)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if authUser.User.Email == "" && authUser.User.Username == "" && authUser.Token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry, user not registered!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Login Successfully!", "data": authUser})
}
