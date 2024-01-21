package main

import (
	"fmt"
	"journey-user/app"
	database "journey-user/app"
	"journey-user/controller"
	"journey-user/repository"
	"journey-user/services"
)

func main() {

	db := database.NewDb()
	userRepository := repository.NewInMemoryUserRepository()
	userService := services.NewUserService(db, userRepository)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(userController)

	fmt.Println(userService)

	router.Run(":8000")
}
