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

// CORSMiddleware enables cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BearerSchema) {
			fmt.Println("Token not found")
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
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
}
