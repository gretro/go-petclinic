package webserver

import (
	"encoding/json"
	"net/http"
	"petclinic-go/server/internal/app/system"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func mapValidationErrors(err validator.ValidationErrors) map[string]string {
	result := make(map[string]string)

	for _, fieldError := range err {
		result[fieldError.Field()] = fieldError.Error()
	}

	return result
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	ginError := c.Errors.Last()
	if ginError != nil {
		switch e := ginError.Err.(type) {
		case *system.EntityNotFoundError:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"title":  "Not Found",
				"detail": e.Error(),
			})
			return
		case *system.EntityConflictError:
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"status": http.StatusConflict,
				"title":  "Conflict",
				"detail": e.Error(),
			})
			return
		case validator.ValidationErrors:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"title":   "Bad Request",
				"detail":  "Validation Error",
				"details": mapValidationErrors(e),
			})
		case *json.UnmarshalTypeError:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"title":  "Bad Request",
				"detail": "Validation Error",
				"details": map[string]string{
					e.Field: e.Error(),
				},
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"title":  "Internal Server Error",
				"detail": e.Error(),
			})
			return
		}
	}
}
