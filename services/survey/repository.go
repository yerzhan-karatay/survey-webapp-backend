package survey

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// Repository includes repository method for survey
type Repository interface {
	CreateSurvey(*models.Survey) error
	CreateQuestionPerSurvey(*models.Question) error
	CreateOptionsPerQuestion(option *models.Option) error
	GetQuestionListBySurveyID(*[]*models.Question, int) error
	GetOptionsByQuestionID(*[]*models.Option, int) error
	GetSurveyListByUserID(*[]*models.Survey, int) error
	GetSurveyByID(*models.Survey, int) error
	UpdateSurvey(*models.Survey) error
	DeleteSurvey(*models.Survey) error
}

type repository struct {
	db *gorm.DB
}

// GetRepository returns Survey Repository
func GetRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// CreateSurvey create survey record
func (r *repository) CreateSurvey(survey *models.Survey) error {
	survey.Created = time.Now()
	r.db.Create(&survey)
	if r.db.Table("survey").Where("id = ?", survey.ID).First(survey).RecordNotFound() {
		return ErrInsertFailed
	}
	return nil
}

// CreateQuestionPerSurvey create question record
func (r *repository) CreateQuestionPerSurvey(question *models.Question) error {
	question.Created = time.Now()
	r.db.Create(&question)
	if r.db.Table("question").Where("id = ?", question.ID).RecordNotFound() {
		return ErrInsertFailed
	}
	return nil
}

// CreateOptionsPerQuestion create options record
func (r *repository) CreateOptionsPerQuestion(option *models.Option) error {
	r.db.Create(&option)
	if r.db.Table("option").Where("id = ?", option.ID).RecordNotFound() {
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

// GetSurveyListByUserID get all surveys by userID
func (r *repository) GetSurveyListByUserID(surveys *[]*models.Survey, userID int) error {
	if err := r.db.Table("survey").Where("user_id = ?", userID).Find(&surveys).Error; err != nil {
		return err
	}
	return nil
}

// GetSurveyByID get survey by id
func (r *repository) GetSurveyByID(survey *models.Survey, surveyID int) error {
	if err := r.db.Table("survey").Where("id = ?", surveyID).Find(survey).Error; err != nil {
		return ErrNotFound
	}
	return nil
}

// GetOptionsByQuestionID get options by question id
func (r *repository) GetOptionsByQuestionID(options *[]*models.Option, questionID int) error {
	if err := r.db.Table("option").Where("question_id = ?", questionID).Find(options).Error; err != nil {
		return ErrNotFound
	}
	return nil
}

// UpdateSurvey update survey by id
func (r *repository) UpdateSurvey(updatedSurvey *models.Survey) error {
	if err := r.db.Table("survey").Save(updatedSurvey).Error; err != nil {
		return ErrUpdateFailed
	}
	return nil
}

// DeleteSurvey delete survey by id
func (r *repository) DeleteSurvey(survey *models.Survey) error {
	r.db.Delete(survey)
	if r.db.Table("survey").Where("id = ?", survey.ID).RecordNotFound() {
		return ErrDeleteFailed
	}
	return nil
}
