package models

import (
	"time"
)

// Option represents the model for a option and is mapping to `option` table
type Option struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string    `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Created         time.Time `json:"created" gorm:"column:created"`
	QuestionID      int       `json:"question_id" gorm:"column:question_id;type:int(11);not null"`
	ResponseAnswers []ResponseAnswer
}

// TableName set the table name to "option"
func (Option) TableName() string {
	return "option"
}
