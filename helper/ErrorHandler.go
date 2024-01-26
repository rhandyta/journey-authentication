package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Message string
	Code    int
}

func NewAppError(errorMessage string, code int) *AppError {
	return &AppError{
		Message: errorMessage,
		Code:    code,
	}
}

func HandleErrorFormatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				HandleError(c, NewAppError("Sorry, something went wrong, please try again later!", http.StatusInternalServerError))
			}

		}()
		c.Next()

		if err, exists := c.Get("Error"); exists {
			if appError, ok := err.(*AppError); ok {
				HandleError(c, appError)
			} else {
				HandleError(c, NewAppError("Sorry, something went wrong, please try again later!", http.StatusInternalServerError))
			}
		}

		delete(c.Keys, "Error")
	}
}

func HandleError(c *gin.Context, err *AppError) {
	c.JSON(err.Code, gin.H{"error_message": err.Message})
	c.Abort()
}
