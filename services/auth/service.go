package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/security"
)

//Service interface
type Service interface {
	Login(ctx *gin.Context, request models.AuthCredentials) (string, error)
}

type service struct {
	authRepository Repository
}

// GetService interface
func GetService(userRepo Repository) Service {
	return &service{
		authRepository: userRepo,
	}
}

// Login godoc
// @Summary User authorization
// @Description User authorization
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param requestBody body models.AuthCredentials true "Login user"
// @Success 200 {object} TokenResponse
// @Failure 400 {string} ErrBadRequest
// @Failure 404 {string} ErrNotFound
// @Router /login [post]
func (s *service) Login(ctx *gin.Context, request models.AuthCredentials) (string, error) {
	loggedInUser, err := s.authRepository.CheckUser(request)
	if err != nil {
		return "", err
	}
	token := security.JWTAuthService().GenerateToken(loggedInUser)
	return token, nil
}

// TokenResponse is the response structure for signup/signin
type TokenResponse struct {
	Token string `json:"token" binding:"required" example:"dummy token"`
}
