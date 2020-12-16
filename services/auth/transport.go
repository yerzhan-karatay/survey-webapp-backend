package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// MakeHTTPHandler mounts auth services to gin handler
func MakeHTTPHandler(r *gin.Engine, s Service) *gin.Engine {
	r.POST("/api/login", func(ctx *gin.Context) {
		var request models.AuthCredentials
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		token, err := s.Login(ctx, request)
		if err != nil {
			ctx.Error(ErrNotFound)
		} else {
			ctx.JSON(http.StatusOK, TokenResponse{
				Token: token,
			})
		}
		return
	})
	return r
}

var (
	// ErrBadRequest means params are not correct
	ErrBadRequest = errors.NewHTTPError(400, "Bad request")
	// ErrNotFound means user was not found in the db or incorrect password
	ErrNotFound = errors.NewHTTPError(http.StatusNotFound, "Not found user")
)
