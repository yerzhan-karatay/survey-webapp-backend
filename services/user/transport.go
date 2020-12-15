package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/security"
)

// MakeHTTPHandler mounts user services to gin handler
func MakeHTTPHandler(r *gin.Engine, s Service) *gin.Engine {
	groupRoutes := r.Group("/api/users")
	groupRoutes.POST("", func(ctx *gin.Context) {
		var request CreateUserRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		token, err := s.CreateUser(ctx, request)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusCreated, TokenResponse{
				Token: token,
			})
		}

		return
	})

	groupRoutes.Use(errors.AuthorizeJWT())
	groupRoutes.GET("/me", func(ctx *gin.Context) {
		userByToken, err := security.JWTAuthService().GetUserByToken(ctx)
		user, err := s.GetUser(ctx, userByToken.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, user)
		}

		return
	})

	return r
}

// TokenResponse is the response structure for user POST api
type TokenResponse struct {
	Token string `json:"token" binding:"required" example:"dummy token"`
}

// CreateUserRequest is the request structure for user POST api
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required" example:"mail@mail.com"`
	Password string `json:"password" binding:"required" example:"123"`
	FullName string `json:"full_name" example:"Yerzhan Karatayev"`
}

var (
	// ErrBadRequest means params are not correct
	ErrBadRequest = errors.NewHTTPError(http.StatusBadRequest, "Bad request")
	// ErrInsertFailed means record is not persusted into table
	ErrInsertFailed = errors.NewHTTPError(http.StatusInternalServerError, "Insert record failed")
	// ErrNotFound means user was not found in the db by userId or email
	ErrNotFound = errors.NewHTTPError(http.StatusNotFound, "User not found")
)
