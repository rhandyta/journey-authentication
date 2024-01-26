package main

import (
	"journey-user/app"
	"journey-user/controller"
	"journey-user/services"
)

func main() {
	db := app.NewDb()
	userService := services.NewUserService(db)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(userController)

	router.Run(":8000")
}
