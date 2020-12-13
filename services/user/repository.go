package user

import (
	"github.com/jinzhu/gorm"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

type Repository interface {
	CreateUser(CreateUserRequest) error
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
func (r *repository) CreateUser(userRequest CreateUserRequest) error {
	user := &models.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		FullName: userRequest.FullName,
	}
	r.db.Create(user)
	if r.db.Table("user").Where("email = ?", user.Email).RecordNotFound() {
		return ErrInsertFailed
	}
	return nil
}
