package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Message string
	Code    int
}

// middleware
func MiddlewareErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				GlobalInternalServerError(c)
			}

		}()
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				switch e.Type {
				case gin.ErrorTypePublic:
					if !c.Writer.Written() {
						c.JSON(c.Writer.Status(), gin.H{"message": e.Error()})
					}
					break
				case gin.ErrorTypeBind:
					if errs, ok := e.Err.(validator.ValidationErrors); ok {
						status := c.Writer.Status()
						if c.Writer.Status() != http.StatusOK {
							status = c.Writer.Status()
						}
						if len(errs) > 0 {
							c.JSON(status, gin.H{"message": ValidationErrorToText(errs[0])})
						}
					} else {
						GlobalInternalServerError(c)
					}
					break
				default:
					GlobalInternalServerError(c)
				}
			}
		}
	}
}

func GlobalInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, Something Went Wrong, Please Try Again Later!"})
	c.Abort()
}

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field(), e.Param())
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}
