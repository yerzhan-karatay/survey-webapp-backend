package services

import (
	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/db"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/user"
)

// MakeHTTPHandler composites service components and mounts them
func MakeHTTPHandler(r *gin.Engine) *gin.Engine {
	db := db.Get()

	// User service
	userRepo := user.GetRepository(db)
	userService := user.GetService(userRepo)

	r = user.MakeHTTPHandler(r, userService)

	return r
}
