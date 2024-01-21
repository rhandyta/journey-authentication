package services

import (
	"database/sql"
	"journey-user/model"
	"journey-user/repository"
	"log"

	"github.com/gin-gonic/gin"
)

type UserServiceImplementation struct {
	UserRepository *repository.InMemoryUserRepository
	db             *sql.DB
}

func NewUserService(db *sql.DB, userRepository *repository.InMemoryUserRepository) *UserServiceImplementation {
	return &UserServiceImplementation{db: db, UserRepository: userRepository}
}

func (userService *UserServiceImplementation) Get(c *gin.Context) []model.UserResponse {

	userRepo := userService.UserRepository
	tx, err := userService.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	users := userRepo.Get(c.Request.Context(), tx)
	var responses []model.UserResponse
	for _, user := range users {
		responses = append(responses, model.UserResponse{
			Id:        user.Id,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Age:       user.Age,
			Email:     user.Email,
		})
	}

	return responses
}
