package survey

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
	groupRoutes := r.Group("/api/surveys")
	groupRoutes.Use(errors.AuthorizeJWT())

	groupRoutes.POST("", func(ctx *gin.Context) {
		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		var request TitleRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		errCreate := s.CreateSurvey(request.Title, user.ID)
		if errCreate != nil {
			ctx.Error(errCreate)
		} else {
			ctx.JSON(http.StatusCreated, gin.H{})
		}

		return
	})

	groupRoutes.GET("", func(ctx *gin.Context) {
		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		surveys, err := s.GetSurveyListByUserID(user.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, surveys)
		}

		return
	})

	groupRoutes.POST("/full", func(ctx *gin.Context) {
		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		var request FullSurveyRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}
		err := s.CreateSurveyWithQnA(request, user.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}

		return
	})

	groupRoutes.GET("/:id/full", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("id"))

		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		surveyWithQnA, err := s.GetSurveyWithQnA(surveyID, user.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, surveyWithQnA)
		}

		return
	})

	groupRoutes.GET("/:id", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("id"))

		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		survey, err := s.GetSurveyByID(surveyID)
		if err != nil {
			ctx.Error(err)
		} else if user.ID != survey.UserID {
			ctx.Error(ErrAccessDenied)
		} else {
			ctx.JSON(http.StatusOK, survey)
		}

		return
	})

	groupRoutes.PUT("/:id", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("id"))

		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		var request TitleRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		err := s.UpdateSurvey(surveyID, request.Title, user.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}

		return
	})

	groupRoutes.DELETE("/:id", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("id"))

		user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		if errToken != nil {
			ctx.Error(errToken)
			return
		}

		err := s.DeleteSurvey(surveyID, user.ID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}

		return
	})

	return r
}

// TitleRequest is the request structure for survey POST api
type TitleRequest struct {
	Title string `json:"title" gorm:"column:title;type:varchar(255);not null;" example:"This is title"`
}

// QuestionRequest is the request structure for survey POST api
type QuestionRequest struct {
	QuestionTitle string   `json:"title" gorm:"column:title;type:varchar(255);not null;" example:"How old are you?"`
	Options       []string `json:"options" gorm:"column:options;type:ARRAY" example:"25,28"`
}

// FullSurveyRequest is the request structure for survey full POST api
type FullSurveyRequest struct {
	SurveyTitle string            `json:"title" gorm:"column:title;type:varchar(255);not null" example:"This is survey title"`
	Questions   []QuestionRequest `json:"questions" gorm:"column:questions;type:ARRAY"`
}

// FullSurveyWithQnAQuestion is the request structure for survey full POST api
type FullSurveyWithQnAQuestion struct {
	QuestionID    int              `json:"id" gorm:"column:id;type:int(11);not null" example:"1"`
	QuestionTitle string           `json:"title" gorm:"column:title;type:varchar(255);not null" example:"This is question Title"`
	Options       []*models.Option `json:"options" gorm:"column:options;type:ARRAY"`
}

// FullSurveyWithQnA is the response structure for survey full GET api
type FullSurveyWithQnA struct {
	Survey    models.Survey               `json:"survey"`
	Questions []FullSurveyWithQnAQuestion `json:"questions" gorm:"column:questions;type:ARRAY"`
}

var (
	// ErrAccessDenied means user don't have access to this survey
	ErrAccessDenied = errors.NewHTTPError(http.StatusForbidden, "Access to this survey denied")
	// ErrBadRequest means params are not correct
	ErrBadRequest = errors.NewHTTPError(http.StatusBadRequest, "Bad request")
	// ErrInsertFailed means record is not persusted into table
	ErrInsertFailed = errors.NewHTTPError(http.StatusInternalServerError, "Insert record failed")
	// ErrNotFound means survey was not found in the db
	ErrNotFound = errors.NewHTTPError(http.StatusNotFound, "Survey not found")
	// ErrUpdateFailed means survey was not updated
	ErrUpdateFailed = errors.NewHTTPError(http.StatusInternalServerError, "Survey update failed")
	// ErrDeleteFailed means survey was not deleted from the db
	ErrDeleteFailed = errors.NewHTTPError(http.StatusInternalServerError, "Survey deletion failed")
)
