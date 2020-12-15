package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
	"github.com/yerzhan-karatay/survey-webapp-backend/models"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/security"
)

// MakeHTTPHandler mounts user services to gin handler
func MakeHTTPHandler(r *gin.Engine, s Service) *gin.Engine {
	groupRoutes := r.Group("/api/survey")
	groupRoutes.Use(errors.AuthorizeJWT())

	groupRoutes.POST("/:surveyID/responses", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		var request []ReponseAnswerRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		errCreate := s.CreateResponse(user.ID, surveyID, request)
		if errCreate != nil {
			ctx.Error(errCreate)
		} else {
			ctx.JSON(http.StatusCreated, gin.H{})
		}

		return
	})

	groupRoutes.GET("/:surveyID/responses", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		responses, err := s.GetResponsesBySurveyID(surveyID, user.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, responses)
		}

		return
	})

	groupRoutes.GET("/:surveyID/responses/:responseID", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		responseID, _ := strconv.Atoi(ctx.Param("responseID"))

		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		responseAns, err := s.GetResponseAnswersByID(responseID, user.ID, surveyID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, responseAns)
		}

		return
	})

	groupRoutesMy := r.Group("/api/responses")
	groupRoutesMy.Use(errors.AuthorizeJWT())

	groupRoutesMy.GET("/my", func(ctx *gin.Context) {
		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		responsedSurveys, err := s.GetResponsedSurveysByUserID(user.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, responsedSurveys)
		}

		return
	})

	return r
}

// ReponseAnswerRequest is the request structure for response POST api
type ReponseAnswerRequest struct {
	QuestionID int `json:"question_id" gorm:"column:title;type:int(11);not null" example:"1"`
	OptionID   int `json:"option_id" gorm:"column:options;type:int(11);not null" example:"1"`
}

// ReponseAnswerResponse is the request structure for response GET apis
type ReponseAnswerResponse struct {
	ResponseAnswers []*models.ResponseAnswer `json:"response_answers" gorm:"column:response_answers;type:array"`
	Response        *models.Response         `json:"response" gorm:"column:response;not null"`
}

var (
	// ErrAccessDenied means user don't have access to this question
	ErrAccessDenied = errors.NewHTTPError(http.StatusForbidden, "Access to this Survey denied")
	// ErrBadRequest means params are not correct
	ErrBadRequest = errors.NewHTTPError(http.StatusBadRequest, "Bad request")
	// ErrInsertFailed means record is not persusted into table
	ErrInsertFailed = errors.NewHTTPError(http.StatusInternalServerError, "Insert record failed")
	// ErrNotFound means Survey was not found in the db
	ErrNotFound = errors.NewHTTPError(http.StatusNotFound, "Survey not found")
	// ErrUpdateFailed means Survey was not updated
	ErrUpdateFailed = errors.NewHTTPError(http.StatusInternalServerError, "Survey update failed")
	// ErrDeleteFailed means Survey was not deleted from the db
	ErrDeleteFailed = errors.NewHTTPError(http.StatusInternalServerError, "Survey deletion failed")
	// ErrAlreadyExist means Survey was not deleted from the db
	ErrAlreadyExist = errors.NewHTTPError(http.StatusConflict, "Survey response already exists")
)
