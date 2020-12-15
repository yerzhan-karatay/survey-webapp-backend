package survey

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
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