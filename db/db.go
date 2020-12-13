package db

import (
	"time"

	"github.com/jinzhu/gorm"
	configs "github.com/yerzhan-karatay/survey-webapp-backend/config"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"

	// DBs
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Connect to the DATABASE
func Connect() (*gorm.DB, error) {
	config := configs.Get()
	db, err := gorm.Open(config.DB.DBtype, config.DB.DBurl)
	if err != nil {
		panic("Attention! Failed to connect database!")
	}

	db.DB().SetMaxIdleConns(config.DB.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DB.MaxOpenConns)

	connMaxLifetime, err := time.ParseDuration(config.DB.ConnMaxLifetime)
	if err != nil {
		panic("Attention! Invalid ConnMaxLifetime!")
	}
	db.DB().SetConnMaxLifetime(connMaxLifetime)

	InitLocalDB(db)

	return db, nil
}

// Get return db object
func Get() *gorm.DB {
	db, err := Connect()
	if err != nil {
		panic("Attention! Failed to connect database!")
	}
	return db
}

// InitLocalDB initialize local database
func InitLocalDB(db *gorm.DB) {
	db.DropTableIfExists("user")

	db.CreateTable(&models.User{})
}
