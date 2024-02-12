package route

import (
	handler "journey-user/internal"
	"journey-user/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.Use(handler.MiddlewareErrorHandle())

	apiGroup := router.Group("/api")

	apiGroup.POST("/auth/sign-up", userController.CreateUser)
	apiGroup.GET("/user/:id", userController.GetUserById)
	apiGroup.GET("/user", userController.GetUserByEmail)

	return router
}
