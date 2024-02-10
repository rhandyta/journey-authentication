package controller

import (
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

	registUser, err := auth.service.Registration(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": registUser})
}

func (auth *AuthenticationControllerImpl) login(c *gin.Context) {
	var userLogin model.Login
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authUser, err := auth.service.Login(userLogin)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if authUser.User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Login Successfully", "user": authUser})
}
