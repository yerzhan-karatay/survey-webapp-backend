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

func (s *service) CreateUser(ctx *gin.Context, request CreateUserRequest) (string, error) {
	var token string = "dummy token"

	err := s.userRepository.CreateUser(request)
	if err != nil {
		return "", err
	}

	return token, nil
}
