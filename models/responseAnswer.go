package models

// ResponseAnswer represents the model for a responseAnswer and is mapping to `responseAnswer` table
type ResponseAnswer struct {
	ID         int `json:"id" gorm:"primaryKey;autoIncrement"`
	ResponseID int `json:"response_id" gorm:"column:response_id;type:int(11);not null"`
	QuestionID int `json:"question_id" gorm:"column:question_id;type:int(11);not null"`
	OptionID   int `json:"option_id" gorm:"column:option_id;type:int(11);not null"`
}

// TableName set the table name to "responseAnswer"
func (ResponseAnswer) TableName() string {
	return "response_answer"
}
