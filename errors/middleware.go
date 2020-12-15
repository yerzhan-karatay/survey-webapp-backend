package errors

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/security"
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

// AuthorizeJWT authorize an access
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := security.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
