package response

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
)

// Repository includes repository method for response
type Repository interface {
	CreateResponse(*models.Response) error
	CreateResponseAnswer(responseAns *models.ResponseAnswer) error
	GetResponseByID(*models.Response, int) error
	GetResponsesByUserID(*[]*models.Response, int) error
	GetResponseCountBySurveyIDnUserID(surveyID int, userID int) (int, error)
	GetResponsesBySurveyID(*[]*models.Response, int) error
	GetResponseAnswersByReponseID(*[]*models.ResponseAnswer, int) error
	GetSurveyByID(survey *models.Survey, surveyID int) error
	GetUserByID(user *models.User, userID int) error
	GetFullResponseByReponseID(quesOpt *[]*QuestionOptionText, responseID int) error
}

type repository struct {
	db *gorm.DB
}

// GetRepository returns Response Repository
func GetRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// CreateResponse create response record
func (r *repository) CreateResponse(response *models.Response) error {
	response.Created = time.Now()
	r.db.Create(response)
	if r.db.Table("response").Where("id = ?", response.ID).RecordNotFound() {
		return ErrInsertFailed
	}
	return nil
}

// CreateResponseAnswer create response answer record
func (r *repository) CreateResponseAnswer(responseAns *models.ResponseAnswer) error {
	r.db.Create(responseAns)
	if r.db.Table("response_answer").Where("id = ?", responseAns.ID).RecordNotFound() {
		return ErrInsertFailed
	}
	return nil
}

// GetResponseByID get response by id
func (r *repository) GetResponseByID(response *models.Response, responseID int) error {
	if err := r.db.Table("response").Where("id = ?", responseID).Find(response).Error; err != nil {
		return err
	}
	return nil
}

// GetResponsesByUserID get responses by user id
func (r *repository) GetResponsesByUserID(responses *[]*models.Response, userID int) error {
	if err := r.db.Table("response").Where("user_id = ?", userID).Find(responses).Error; err != nil {
		return err
	}
	return nil
}

// GetResponsesBySurveyID get responses by survey id
func (r *repository) GetResponseCountBySurveyIDnUserID(surveyID int, userID int) (int, error) {
	var responses int
	if err := r.db.Table("response").Where("survey_id = ? AND user_id = ?", surveyID, userID).Count(&responses).Error; err != nil {
		return responses, err
	}
	return responses, nil
}

// GetResponsesBySurveyID get responses by survey id
func (r *repository) GetResponsesBySurveyID(responses *[]*models.Response, surveyID int) error {
	if err := r.db.Table("response").Where("survey_id = ?", surveyID).Find(responses).Error; err != nil {
		return err
	}
	return nil
}

// GetResponseAnswersByReponseID get response answers by responseID
func (r *repository) GetResponseAnswersByReponseID(responseAns *[]*models.ResponseAnswer, responseID int) error {
	if err := r.db.Table("response_answer").Where("response_id = ?", responseID).Find(responseAns).Error; err != nil {
		return err
	}
	return nil
}

// GetFullResponseByReponseID get response answers by responseID
func (r *repository) GetFullResponseByReponseID(quesOpt *[]*QuestionOptionText, responseID int) error {
	if err := r.db.Table("response_answer").Select("option.title as option, question.title as question").Where("response_answer.response_id = ?", responseID).Joins("JOIN option on option.id = response_answer.option_id").Joins("JOIN question on question.id = response_answer.question_id").Find(quesOpt).Error; err != nil {
		return err
	}
	return nil
}

// GetSurveyByID get survey by id
func (r *repository) GetSurveyByID(survey *models.Survey, surveyID int) error {
	if err := r.db.Table("survey").Where("id = ?", surveyID).Find(survey).Error; err != nil {
		return err
	}
	return nil
}

// GetUser get user record by id
func (r *repository) GetUserByID(user *models.User, userID int) error {
	if r.db.Table("user").Select("email").Where("id = ?", userID).First(&user).RecordNotFound() {
		return ErrNotFound
	}
	return nil
}
