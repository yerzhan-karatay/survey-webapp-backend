package errors

import (
	"github.com/gin-gonic/gin"
)

// HandleHTTPError returns http code and error if error exist in the context
func HandleHTTPError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		if httpErr, ok := err.Err.(HTTPError); ok {
			c.JSON(httpErr.GetStatusCode(), gin.H{
				"error": httpErr.GetMessage(),
			})
			return
		}

		c.JSON(500, gin.H{
			"error": err.Err.Error(),
		})
		return
	}
}
