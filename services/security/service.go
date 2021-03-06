package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/config"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

//JWTService includes generate and validate methods
type JWTService interface {
	GenerateToken(userInfo models.User) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserByToken(c *gin.Context) (models.User, error)
}
type authCustomClaims struct {
	Email    string `json:"email"`
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//JWTAuthService auth service
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Yerzhan",
	}
}

func getSecretKey() string {
	cfg := config.Get()
	secret := cfg.SECURITY.SecretKey
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (s *jwtServices) GenerateToken(userInfo models.User) string {
	claims := &authCustomClaims{
		userInfo.Email,
		userInfo.ID,
		userInfo.FullName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])

		}
		return []byte(s.secretKey), nil
	})

}

func (s *jwtServices) GetUserByToken(c *gin.Context) (models.User, error) {
	claims := &authCustomClaims{}
	user := models.User{}

	const BearerSchema = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if token != nil {
		user.ID = claims.ID
		user.Email = claims.Email
		user.FullName = claims.FullName
		return user, nil
	}
	return user, err
}
