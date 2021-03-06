package user

import (
	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/security"
)

// Service is the interface of User service
type Service interface {
	CreateUser(ctx *gin.Context, request CreateUserRequest) (string, error)
	GetUser(ctx *gin.Context, userID int) (models.User, error)
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
// @Success 201 {object} TokenResponse
// @Failure 400 {string} ErrBadRequest
// @Failure 403 {string} ErrUserExists
// @Failure 500 {string} ErrInsertFailed
// @Router /api/users [post]
func (s *service) CreateUser(ctx *gin.Context, request CreateUserRequest) (string, error) {
	_, errUserCheck := s.userRepository.GetUserByEmail(request.Email)
	if errUserCheck != ErrNotFound {
		return "", ErrUserExists
	}

	user, err := s.userRepository.CreateUser(request)
	if err != nil {
		return "", err
	}
	token := security.JWTAuthService().GenerateToken(user)
	return token, nil
}

// GetUser godoc
// @Summary get user info
// @Description User information by token
// @Security ApiKeyAuth
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 404 {string} ErrNotFound
// @Router /api/users/me [get]
func (s *service) GetUser(ctx *gin.Context, userID int) (models.User, error) {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}
