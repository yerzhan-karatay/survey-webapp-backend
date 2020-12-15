package survey

import (
	"github.com/jinzhu/gorm"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// Repository includes repository method for survey
type Repository interface {
	CreateSurvey(*models.Survey) error
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
	r.db.Create(survey)
	if r.db.Table("survey").Where("id = ?", survey.ID).RecordNotFound() {
		return ErrInsertFailed
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
