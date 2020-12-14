package errors

import (
	"github.com/gin-gonic/gin"
)

// ErrorStructure is the response structure for error messages
type ErrorStructure struct {
	Error string `json:"error" binding:"required" example:"error message"`
}

// HandleHTTPError returns http code and error if error exist in the context
func HandleHTTPError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		if httpErr, ok := err.Err.(HTTPError); ok {
			c.JSON(httpErr.GetStatusCode(), ErrorStructure{
				Error: httpErr.GetMessage(),
			})
			return
		}

		c.JSON(500, ErrorStructure{
			Error: err.Err.Error(),
		})
		return
	}
}
