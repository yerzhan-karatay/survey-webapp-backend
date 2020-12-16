package question

import (
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// Service is the interface of Question service
type Service interface {
	CreateQuestion(string, int) error
	GetQuestionListBySurveyID(int) ([]*models.Question, error)
	GetQuestionByID(int) (models.Question, error)
	UpdateQuestion(int, string, int) error
	DeleteQuestion(int, int) error
}

type service struct {
	QuestionRepository Repository
}

// GetService returns Question service
func GetService(QuestionRepo Repository) Service {
	return &service{
		QuestionRepository: QuestionRepo,
	}
}

// CreateQuestion godoc
// @Summary Add a new Question
// @Description Question creation
// @Security ApiKeyAuth
// @Tags Questions
// @Accept  json
// @Produce  json
// @Param surveyID path int true "Survey ID"
// @Param requestBody body TitleRequest true "Question title"
// @Success 201
// @Failure 400 {string} ErrBadRequest
// @Failure 500 {string} ErrInsertFailed
// @Router /api/survey/{surveyID}/questions [post]
func (s *service) CreateQuestion(questionTitle string, surveyID int) error {
	var survey models.Survey
	errSurvey := s.QuestionRepository.GetSurveyByID(&survey, surveyID)
	if errSurvey != nil {
		return errSurvey
	}

	if questionTitle == "" {
		return ErrBadRequestTitle
	}

	question := &models.Question{
		Title:    questionTitle,
		SurveyID: surveyID,
	}
	err := s.QuestionRepository.CreateQuestion(question)
	if err != nil {
		return err
	}
	return nil
}

// GetQuestionListBySurveyID godoc
// @Summary Get Question list by survey ID
// @Description Question list by survey ID
// @Security ApiKeyAuth
// @Tags Questions
// @Accept  json
// @Produce  json
// @Param surveyID path int true "Survey ID"
// @Success 200 {array} models.Question
// @Failure 404 {string} ErrNotFound
// @Router /api/survey/{surveyID}/questions [get]
func (s *service) GetQuestionListBySurveyID(surveyID int) ([]*models.Question, error) {
	var survey models.Survey
	errSurvey := s.QuestionRepository.GetSurveyByID(&survey, surveyID)
	if errSurvey != nil {
		return nil, errSurvey
	}

	var questions []*models.Question
	err := s.QuestionRepository.GetQuestionListBySurveyID(&questions, surveyID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// GetQuestionByID godoc
// @Summary Get Question info by ID
// @Description Question information by ID
// @Security ApiKeyAuth
// @Tags Questions
// @Accept  json
// @Produce  json
// @Param surveyID path int true "Survey ID"
// @Param questionID path int true "Question ID"
// @Success 200 {object} models.Question
// @Failure 404 {string} ErrNotFound
// @Failure 403 {string} ErrAccessDenied
// @Router /api/survey/{surveyID}/questions/{questionID} [get]
func (s *service) GetQuestionByID(questionID int) (models.Question, error) {
	var question models.Question

	err := s.QuestionRepository.GetQuestionByID(&question, questionID)
	if err != nil {
		return question, err
	}
	return question, nil
}

// UpdateQuestion godoc
// @Summary Update Question by ID
// @Description Update Question by ID
// @Security ApiKeyAuth
// @Tags Questions
// @Accept  json
// @Produce  json
// @Param requestBody body TitleRequest true "Question title"
// @Param surveyID path int true "Survey ID"
// @Param questionID path int true "Question ID"
// @Success 204
// @Failure 404 {string} ErrNotFound
// @Failure 403 {string} ErrAccessDenied
// @Router /api/survey/{surveyID}/questions/{questionID} [put]
func (s *service) UpdateQuestion(questionID int, newTitle string, surveyID int) error {
	var survey models.Survey
	errSurvey := s.QuestionRepository.GetSurveyByID(&survey, surveyID)
	if errSurvey != nil {
		return errSurvey
	}

	var question models.Question

	err := s.QuestionRepository.GetQuestionByID(&question, questionID)
	if err != nil {
		return err
	}

	if surveyID != question.SurveyID {
		return ErrAccessDenied
	}

	if newTitle == "" {
		return ErrBadRequestTitle
	}

	question.Title = newTitle

	errUpdate := s.QuestionRepository.UpdateQuestion(&question)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

// DeleteQuestion godoc
// @Summary Delete Question by ID
// @Description Delete Question by ID
// @Security ApiKeyAuth
// @Tags Questions
// @Accept  json
// @Produce  json
// @Param surveyID path int true "Survey ID"
// @Param questionID path int true "Question ID"
// @Success 204
// @Failure 404 {string} ErrNotFound
// @Failure 403 {string} ErrAccessDenied
// @Router /api/survey/{surveyID}/questions/{questionID} [delete]
func (s *service) DeleteQuestion(questionID int, surveyID int) error {
	var survey models.Survey
	errSurvey := s.QuestionRepository.GetSurveyByID(&survey, surveyID)
	if errSurvey != nil {
		return errSurvey
	}

	var question models.Question

	err := s.QuestionRepository.GetQuestionByID(&question, questionID)
	if err != nil {
		return err
	}

	if surveyID != question.SurveyID {
		return ErrAccessDenied
	}

	errDelete := s.QuestionRepository.DeleteQuestion(&question)
	if errDelete != nil {
		return errDelete
	}
	return nil
}
