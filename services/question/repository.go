package question

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// Repository includes repository method for question
type Repository interface {
	CreateQuestion(*models.Question) error
	GetQuestionListBySurveyID(*[]*models.Question, int) error
	GetQuestionByID(*models.Question, int) error
	UpdateQuestion(*models.Question) error
	DeleteQuestion(*models.Question) error
	GetSurveyByID(*models.Survey, int) error
}

type repository struct {
	db *gorm.DB
}

// GetRepository returns Question Repository
func GetRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// CreateQuestion create question record
func (r *repository) CreateQuestion(question *models.Question) error {
	question.Created = time.Now()
	r.db.Create(question)
	if r.db.Table("question").Where("id = ?", question.ID).RecordNotFound() {
		return ErrInsertFailed
	}
	return nil
}

// GetQuestionListBySurveyID get all questions by surveyID
func (r *repository) GetQuestionListBySurveyID(questions *[]*models.Question, surveyID int) error {
	if err := r.db.Table("question").Where("survey_id = ?", surveyID).Find(&questions).Error; err != nil {
		return err
	}
	return nil
}

// GetQuestionByID get question by id
func (r *repository) GetQuestionByID(question *models.Question, questionID int) error {
	if err := r.db.Table("question").Where("id = ?", questionID).Find(question).Error; err != nil {
		return ErrNotFound
	}
	return nil
}

// UpdateQuestion update question by id
func (r *repository) UpdateQuestion(updatedQuestion *models.Question) error {
	if err := r.db.Table("question").Save(updatedQuestion).Error; err != nil {
		return ErrUpdateFailed
	}
	return nil
}

// DeleteQuestion delete question by id
func (r *repository) DeleteQuestion(question *models.Question) error {
	r.db.Delete(question)
	if r.db.Table("question").Where("id = ?", question.ID).RecordNotFound() {
		return ErrDeleteFailed
	}
	return nil
}

// GetSurveyByID get survey by id
func (r *repository) GetSurveyByID(survey *models.Survey, surveyID int) error {
	if err := r.db.Table("survey").Where("id = ?", surveyID).Find(survey).Error; err != nil {
		return ErrNotFoundSurvey
	}
	return nil
}
