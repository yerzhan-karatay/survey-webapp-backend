package user

import (
	"github.com/gin-gonic/gin"
)

// Service is the interface of User service
type Service interface {
	CreateUser(ctx *gin.Context, request CreateUserRequest) (string, error)
}

type service struct {
	userRepository Repository
}

// GetService returns User service
func GetService(userRepo Repository) Service {
	return &service{
		userRepository: userRepo,
	}
}

// CreateUser godoc
// @Summary Add a new user
// @Description User creation
// @Tags Users
// @Accept  json
// @Produce  json
// @Param requestBody body CreateUserRequest true "Create user"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {string} ErrBadRequest
// @Failure 500 {string} ErrInsertFailed
// @Router /users [post]
func (s *service) CreateUser(ctx *gin.Context, request CreateUserRequest) (string, error) {
	var token string = "dummy token"

	err := s.userRepository.CreateUser(request)
	if err != nil {
		return "", err
	}

	return token, nil
}
