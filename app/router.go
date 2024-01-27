package app

import (
	"journey-user/controller"
	handler "journey-user/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserControllerImplementation) *gin.Engine {

	route := gin.Default()
	route.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK!"})
	})

	route.Use(handler.MiddlewareErrorHandle())

	apiGroup := route.Group("/api")
	users := apiGroup.Group("/users")
	users.GET("/", userController.Get)
	authentication := apiGroup.Group("/authentication")
	authentication.POST("/registration", userController.Registration)
	return route
}
