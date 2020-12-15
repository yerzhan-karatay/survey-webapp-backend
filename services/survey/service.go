package survey

import (
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// Service is the interface of Survey service
type Service interface {
	CreateSurvey(surveyTitle string, userID int) error
	GetSurveyListByUserID(int) ([]*models.Survey, error)
	GetSurveyByID(int) (models.Survey, error)
	UpdateSurvey(int, string, int) error
	DeleteSurvey(int, int) error
}

type service struct {
	SurveyRepository Repository
}

// GetService returns Survey service
func GetService(SurveyRepo Repository) Service {
	return &service{
		SurveyRepository: SurveyRepo,
	}
}

// CreateSurvey godoc
// @Summary Add a new Survey
// @Description Survey creation
// @Security ApiKeyAuth
// @Tags Surveys
// @Accept  json
// @Produce  json
// @Param requestBody body TitleRequest true "Survey title"
// @Success 201
// @Failure 400 {string} ErrBadRequest
// @Failure 500 {string} ErrInsertFailed
// @Router /api/surveys [post]
func (s *service) CreateSurvey(surveyTitle string, userID int) error {
	survey := &models.Survey{
		Title:  surveyTitle,
		UserID: userID,
	}
	err := s.SurveyRepository.CreateSurvey(survey)
	if err != nil {
		return err
	}
	return nil
}

// GetSurveyListByUserID godoc
// @Summary Get Survey list by user ID
// @Description Survey list for loggedin user
// @Security ApiKeyAuth
// @Tags Surveys
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Survey
// @Failure 404 {string} ErrNotFound
// @Router /api/surveys [get]
func (s *service) GetSurveyListByUserID(surveyID int) ([]*models.Survey, error) {
	var surveys []*models.Survey
	err := s.SurveyRepository.GetSurveyListByUserID(&surveys, surveyID)
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveyByID godoc
// @Summary Get Survey info by ID
// @Description Survey information by ID
// @Security ApiKeyAuth
// @Tags Surveys
// @Accept  json
// @Produce  json
// @Param id path int true "Survey ID"
// @Success 200 {object} models.Survey
// @Failure 404 {string} ErrNotFound
// @Failure 403 {string} ErrAccessDenied
// @Router /api/surveys/{id} [get]
func (s *service) GetSurveyByID(surveyID int) (models.Survey, error) {
	var survey models.Survey

	err := s.SurveyRepository.GetSurveyByID(&survey, surveyID)
	if err != nil {
		return survey, err
	}
	return survey, nil
}

// UpdateSurvey godoc
// @Summary Update Survey by ID
// @Description Update Survey by ID
// @Security ApiKeyAuth
// @Tags Surveys
// @Accept  json
// @Produce  json
// @Param requestBody body TitleRequest true "Survey title"
// @Param id path int true "Survey ID"
// @Success 204 {object} models.Survey
// @Failure 404 {string} ErrNotFound
// @Failure 403 {string} ErrAccessDenied
// @Router /api/surveys/{id} [put]
func (s *service) UpdateSurvey(surveyID int, newTitle string, userID int) error {
	var survey models.Survey

	err := s.SurveyRepository.GetSurveyByID(&survey, surveyID)
	if err != nil {
		return err
	}

	if userID != survey.UserID {
		return ErrAccessDenied
	}

	survey.Title = newTitle

	errUpdate := s.SurveyRepository.UpdateSurvey(&survey)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

// DeleteSurvey godoc
// @Summary Delete Survey by ID
// @Description Delete Survey by ID
// @Security ApiKeyAuth
// @Tags Surveys
// @Accept  json
// @Produce  json
// @Param id path int true "Survey ID"
// @Success 204 {object} models.Survey
// @Failure 404 {string} ErrNotFound
// @Failure 403 {string} ErrAccessDenied
// @Router /api/surveys/{id} [delete]
func (s *service) DeleteSurvey(surveyID int, userID int) error {
	var survey models.Survey

	err := s.SurveyRepository.GetSurveyByID(&survey, surveyID)
	if err != nil {
		return err
	}

	if userID != survey.UserID {
		return ErrAccessDenied
	}

	errDelete := s.SurveyRepository.DeleteSurvey(&survey)
	if errDelete != nil {
		return errDelete
	}
	return nil
}
