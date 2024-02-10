package journey_middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
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
			e := c.Errors[0]
			// gorm validation unique
			if gormErr, ok := e.Err.(*pgconn.PgError); ok {
				// unique code sql
				if gormErr.Code == "23505" {
					messageReplace := strings.ReplaceAll(gormErr.ConstraintName, fmt.Sprintf("idx_%s_", gormErr.TableName), "")
					c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("Sorry, %s already registered!", messageReplace)})
					return
				}
			}

			// if errs, ok := e.Err.(validator.ValidationErrors); ok {
			// 	status := c.Writer.Status()
			// 	if c.Writer.Status() != http.StatusOK {
			// 		status = c.Writer.Status()
			// 	}
			// 	if len(errs) > 0 {
			// 		c.JSON(status, gin.H{"message": ValidationErrorToText(errs[0])})
			// 	}
			// } else {
			// 	GlobalInternalServerError(c)
			// }
		}
	}
}

func GlobalInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, Something Went Wrong, Please Try Again Later!"})
	c.Abort()
}
