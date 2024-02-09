package controllers

import (
	models "journey-user/pkg/models"
	"journey-user/pkg/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		if strings.Contains(err.Error(), "user with the same email already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully", "user": user})
}

func (uc *UserController) GetUserById(c *gin.Context) {
	userID := c.Param("id")
	userIdToInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	user, err := uc.userService.GetUserById(uint(userIdToInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")

	user, err := uc.userService.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
