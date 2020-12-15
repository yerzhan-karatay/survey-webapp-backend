package auth

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
	"github.com/yerzhan-karatay/survey-webapp-backend/utils"
)

// Repository contains check user for login
type Repository interface {
	CheckUser(models.AuthCredentials) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

// GetRepository returns User Repository
func GetRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// CheckUser create user record
func (r *repository) CheckUser(authRequest models.AuthCredentials) (models.User, error) {
	user := &models.AuthCredentials{
		Email:    authRequest.Email,
		Password: authRequest.Password,
	}
	loggedInUser := &models.User{}
	if r.db.Table("user").Where("email = ?", user.Email).First(&loggedInUser).RecordNotFound() {
		log.Println("Login - not found email:", user.Email)
		return *loggedInUser, ErrNotFound
	}
	err := utils.VerifyPassword(loggedInUser.Password, user.Password)
	if err != nil {
		log.Println("Login - password verification failed:", user.Email)
		return *loggedInUser, ErrNotFound
	}
	return *loggedInUser, nil
}
