package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
)

// MakeHTTPHandler mounts user services to gin handler
func MakeHTTPHandler(r *gin.Engine, s Service) *gin.Engine {
	r.POST("/users", func(ctx *gin.Context) {
		var request CreateUserRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		token, err := s.CreateUser(ctx, request)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusCreated, gin.H{
				"token": token,
			})
		}

		return
	})

	return r
}

// CreateUserRequest is the request structure for user POST api
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"`
}

var (
	// ErrBadRequest means params are not correct
	ErrBadRequest = errors.NewHTTPError(400, "bad request")
	// ErrInsertFailed means record is not persusted into table
	ErrInsertFailed = errors.NewHTTPError(500, "insert record failed")
)
