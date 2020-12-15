package user

import (
	"github.com/jinzhu/gorm"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// Repository includes repository method for user
type Repository interface {
	CreateUser(CreateUserRequest) (models.User, error)
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

// CreateUser create user record
func (r *repository) CreateUser(userRequest CreateUserRequest) (models.User, error) {
	user := &models.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		FullName: userRequest.FullName,
	}
	r.db.Create(user)
	if r.db.Table("user").Where("email = ?", user.Email).RecordNotFound() {
		return *user, ErrInsertFailed
	}
	return *user, nil
}
