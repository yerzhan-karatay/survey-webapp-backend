package models

import "time"

// Question represents the model for a question and is mapping to `question` table
type Question struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string    `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Created         time.Time `json:"created" gorm:"column:created"`
	SurveyID        int       `json:"survey_id" gorm:"column:survey_id;type:int(11);not null"`
	Options         []Option
	ResponseAnswers []ResponseAnswer
}

// TableName set the table name to "question"
func (Question) TableName() string {
	return "question"
}
