package models

import (
	"time"
)

// Survey represents the model for a survey and is mapping to `survey` table
type Survey struct {
	ID      int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title   string    `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Created time.Time `json:"created" gorm:"autoCreateTime;column:created;type:timestamp"`
	UserID  int       `json:"user_id" gorm:"column:user_id;type:int(11);not null"`
}

// TableName set the table name to "survey"
func (Survey) TableName() string {
	return "survey"
}
