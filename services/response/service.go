package response

import (
	"fmt"
	"log"

	configs "github.com/yerzhan-karatay/survey-webapp-backend/config"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
	"github.com/yerzhan-karatay/survey-webapp-backend/utils"
)

// Service is the interface of Response service
type Service interface {
	CreateResponse(int, int, []ReponseAnswerRequest) error
	GetResponsesBySurveyID(int, int) ([]ReponseAnswerResponse, error)
	GetResponsedSurveysByUserID(int) ([]RespondedSurveys, error)
	GetResponseAnswersByID(int, int, int) ([]*models.ResponseAnswer, error)
}

type service struct {
	ResponseRepository Repository
}

// GetService returns Response service
func GetService(ResponseRepo Repository) Service {
	return &service{
		ResponseRepository: ResponseRepo,
	}
}

// CreateResponse godoc
// @Summary Add a new Response with answers
// @Description Response creation
// @Security ApiKeyAuth
// @Tags Responses
// @Accept  json
// @Produce  json
// @Param surveyID path int true "Survey ID"
// @Param requestBody body []ReponseAnswerRequest true "Request body"
// @Success 201
// @Failure 400 {string} ErrBadRequest
// @Failure 500 {string} ErrInsertFailed
// @Router /api/survey/{surveyID}/responses [post]
func (s *service) CreateResponse(userID int, surveyID int, responseAns []ReponseAnswerRequest) error {
	// Step 1 - Validate existance and Create response
	resCount, errCount := s.ResponseRepository.GetResponseCountBySurveyIDnUserID(surveyID, userID)
	if errCount != nil {
		return errCount
	}

	if resCount > 0 {
		return ErrAlreadyExist
	}

	var survey models.Survey
	errServ := s.ResponseRepository.GetSurveyByID(&survey, surveyID)
	if errServ != nil {
		return ErrNotFound
	}

	response := &models.Response{
		UserID:   userID,
		SurveyID: surveyID,
	}
	err := s.ResponseRepository.CreateResponse(response)
	if err != nil {
		return err
	}

	// Step 2 - Create response answers
	for _, res := range responseAns {
		resAns := &models.ResponseAnswer{
			ResponseID: response.ID,
			QuestionID: res.QuestionID,
			OptionID:   res.OptionID,
		}

		err := s.ResponseRepository.CreateResponseAnswer(resAns)
		if err != nil {
			return err
		}
	}

	// SEND EMAIL using SMTP
	config := configs.Get()
	if config.SMTP.Email != "CHANGE_TO_YOUR_EMAIL@gmail.com" {
		// GET USER EMAIL
		user := &models.User{}
		s.ResponseRepository.GetUserByID(user, userID)
		emails := []string{user.Email}

		// GET QUESTION OPTION TEXT
		quesOpt := []*QuestionOptionText{}
		errQO := s.ResponseRepository.GetFullResponseByReponseID(&quesOpt, response.ID)
		if errQO != nil {
			log.Println("SMTP - sql query error -", errQO)
		}
		var emailMessage string
		emailMessage = fmt.Sprintf("Survey name - %s\n\n", survey.Title)
		for _, questionAndOption := range quesOpt {
			log.Println("SMTP - RESPONSE -", questionAndOption.Question, questionAndOption.Option)
			emailMessage = fmt.Sprintf("%s\nQuestion - %s\nOption - %s", emailMessage, questionAndOption.Question, questionAndOption.Option)
		}
		utils.SendEmailSMTP(emails, emailMessage)
	}
	return nil
}

// GetResponsesByUserID godoc
// @Summary Get Responded Survey list by user ID
// @Description Responded Survey list by user ID
// @Security ApiKeyAuth
// @Tags Responses
// @Accept  json
// @Produce  json
// @Success 200 {array} RespondedSurveys
// @Failure 403 {string} ErrAccessDenied
// @Failure 404 {string} ErrNotFound
// @Router /api/responses/my [get]
func (s *service) GetResponsedSurveysByUserID(userID int) ([]RespondedSurveys, error) {
	var responses []*models.Response
	err := s.ResponseRepository.GetResponsesByUserID(&responses, userID)
	if err != nil {
		return nil, err
	}

	if len(responses) == 0 {
		surveys := make([]RespondedSurveys, 0)
		return surveys, nil
	}

	respondedSurveyList := make([]RespondedSurveys, len(responses)-1)
	for _, response := range responses {

		var survey models.Survey
		errSurvey := s.ResponseRepository.GetSurveyByID(&survey, response.SurveyID)
		if errSurvey != nil {
			return nil, errSurvey
		}
		newResSur := RespondedSurveys{
			SurveyID:    survey.ID,
			ResponseID:  response.ID,
			SurveyTitle: survey.Title,
			Created:     response.Created,
		}
		respondedSurveyList = append(respondedSurveyList, newResSur)
	}

	return respondedSurveyList, nil
}

// GetResponseAnswersByID godoc
// @Summary Get Response answers by response ID
// @Description Response answers by response ID
// @Security ApiKeyAuth
// @Tags Responses
// @Accept  json
// @Produce  json
// @Param surveyID path int true "Survey ID"
// @Param responseID path int true "Response ID"
// @Success 200 {object} []models.ResponseAnswer
// @Failure 404 {string} ErrNotFound
// @Failure 403 {string} ErrAccessDenied
// @Router /api/survey/{surveyID}/responses/{responseID} [get]
func (s *service) GetResponseAnswersByID(responseID int, userID int, surveyID int) ([]*models.ResponseAnswer, error) {
	var response models.Response
	var responseAns []*models.ResponseAnswer

	err := s.ResponseRepository.GetResponseByID(&response, responseID)
	if err != nil {
		return responseAns, ErrNotFound
	}

	if response.SurveyID != surveyID {
		return responseAns, ErrNotFound
	}

	if response.UserID != userID {
		return responseAns, ErrAccessDenied
	}

	errRespAns := s.ResponseRepository.GetResponseAnswersByReponseID(&responseAns, responseID)
	if errRespAns != nil {
		return nil, errRespAns
	}
	return responseAns, nil
}

// GetResponsesBySurveyID godoc
// @Summary Get Response list by survey ID
// @Description Response list by survey ID
// @Security ApiKeyAuth
// @Tags Responses
// @Accept  json
// @Produce  json
// @Param surveyID path int true "Survey ID"
// @Success 200 {array} ReponseAnswerResponse
// @Failure 403 {string} ErrAccessDenied
// @Failure 404 {string} ErrNotFound
// @Router /api/survey/{surveyID}/responses [get]
func (s *service) GetResponsesBySurveyID(userID int, surveyID int) ([]ReponseAnswerResponse, error) {
	// Step 1 - validate owner of this survey
	var survey models.Survey
	errSurvey := s.ResponseRepository.GetSurveyByID(&survey, surveyID)
	if errSurvey != nil {
		return nil, errSurvey
	}

	if survey.UserID != userID {
		return nil, ErrAccessDenied
	}

	// Step 2 - get all responses by survey
	var responses []*models.Response
	errResp := s.ResponseRepository.GetResponsesBySurveyID(&responses, surveyID)
	if errResp != nil {
		return nil, errResp
	}

	// Step 3 - get all answers per response
	respAnsListForResponse := make([]ReponseAnswerResponse, len(responses)-1)
	for _, response := range responses {
		var respAns []*models.ResponseAnswer

		s.ResponseRepository.GetResponseAnswersByReponseID(&respAns, response.ID)

		fullResponse := ReponseAnswerResponse{
			Response:        response,
			ResponseAnswers: respAns,
		}
		respAnsListForResponse = append(respAnsListForResponse, fullResponse)
	}

	return respAnsListForResponse, nil
}
