package main

import (
	"journey-user/app"
	"journey-user/pkg/controllers"
	"journey-user/pkg/repositories"
	"journey-user/pkg/route"
	"journey-user/pkg/services"
)

func main() {
	db := app.NewDb()
	// userService := services.NewUserService(db)
	// userController := controller.NewUserController(userService)

	// router := app.NewRouter(userController)

	// router.Run(":8000")

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router := route.SetupRouter(userController)

	router.Run(":8000")
}
