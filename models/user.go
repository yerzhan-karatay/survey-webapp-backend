package models

import "time"

// User represents the model for an user and is mapping to `user` table
type User struct {
	ID       int       `json:"id" gorm:"column:id;primary_key;type:int(11);autoIncrement"`
	Email    string    `json:"email" gorm:"column:email;type:varchar(255);not null;unique"`
	Password string    `json:"password" gorm:"column:password;type:varchar(100);not null"`
	FullName string    `json:"full_name" gorm:"column:full_name;type:varchar(100);not null"`
	Created  time.Time `json:"created" gorm:"column:created;type:timestamp;autoCreateTime"`
}

// TableName set the table name to "user"
func (User) TableName() string {
	return "user"
}
