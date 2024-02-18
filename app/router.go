package app

import (
	"journey-user/controller"
	journey_middleware "journey-user/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserControllerImpl, authController *controller.AuthenticationControllerImpl) *gin.Engine {

	route := gin.Default()
	route.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK!"})
	})

	route.Use(journey_middleware.MiddlewareErrorHandle())

	apiGroup := route.Group("/api")
	users := apiGroup.Group("/users")
	authentication := apiGroup.Group("/authentication")

	authentication.POST("/registration", authController.Registration)
	authentication.POST("/login", authController.Login)

	users.GET("/", userController.Get)
	return route

}
