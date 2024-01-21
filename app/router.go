package app

import (
	"journey-user/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserControllerImplementation) *gin.Engine {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK!"})
	})

	apiGroup := r.Group("/api")
	user := apiGroup.Group("/user")
	user.GET("/", userController.Get)
	user.POST("/registration", userController.Get)
	// authentication := apiGroup.Group("/auth")
	// authentication.GET("registration", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "API REGIST!"})
	// })

	return r
}
