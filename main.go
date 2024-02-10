package main

import (
	"journey-user/app"
	"journey-user/controller"
)

func main() {
	db := app.NewDb()
	userController := controller.NewUserController(db)
	authController := controller.NewAuthenticationController(db)
	router := app.NewRouter(userController, authController)

	router.Run(":8000")
}
