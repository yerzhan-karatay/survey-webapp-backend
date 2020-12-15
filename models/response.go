package models

import (
	"time"
)

// Response represents the model for a response and is mapping to `response` table
type Response struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Created         time.Time `json:"created" gorm:"column:created"`
	QuestionID      int       `json:"question_id" gorm:"column:question_id;type:int(11);not null"`
	UserID          int       `json:"user_id" gorm:"column:user_id;type:int(11);not null"`
	ResponseAnswers []ResponseAnswer
}

// TableName set the table name to "response"
func (Response) TableName() string {
	return "response"
}
