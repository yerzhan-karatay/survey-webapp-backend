package question

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
)

// MakeHTTPHandler mounts user services to gin handler
func MakeHTTPHandler(r *gin.Engine, s Service) *gin.Engine {
	// TODO: check if user is owner of the question's survey
	groupRoutes := r.Group("/api/survey")
	groupRoutes.Use(errors.AuthorizeJWT())

	groupRoutes.POST("/:surveyID/questions", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		// user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		// if errToken != nil {
		// 	ctx.Error(errToken)
		// 	return
		// }

		var request TitleRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		errCreate := s.CreateQuestion(request.Title, surveyID)
		if errCreate != nil {
			ctx.Error(errCreate)
		} else {
			ctx.JSON(http.StatusCreated, gin.H{})
		}

		return
	})

	groupRoutes.GET("/:surveyID/questions", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		// user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		// if errToken != nil {
		// 	ctx.Error(errToken)
		// 	return
		// }

		questions, err := s.GetQuestionListBySurveyID(surveyID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, questions)
		}

		return
	})

	groupRoutes.GET("/:surveyID/questions/:questionID", func(ctx *gin.Context) {
		// surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		questionID, _ := strconv.Atoi(ctx.Param("questionID"))

		// user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		// if errToken != nil {
		// 	ctx.Error(errToken)
		// 	return
		// }

		question, err := s.GetQuestionByID(questionID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, question)
		}

		return
	})

	groupRoutes.PUT("/:surveyID/questions/:questionID", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		questionID, _ := strconv.Atoi(ctx.Param("questionID"))

		// user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		// if errToken != nil {
		// 	ctx.Error(errToken)
		// 	return
		// }

		var request TitleRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(ErrBadRequest)
			return
		}

		err := s.UpdateQuestion(questionID, request.Title, surveyID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}

		return
	})

	groupRoutes.DELETE("/:surveyID/questions/:questionID", func(ctx *gin.Context) {
		surveyID, _ := strconv.Atoi(ctx.Param("surveyID"))
		questionID, _ := strconv.Atoi(ctx.Param("questionID"))

		// user, errToken := security.JWTAuthService().GetUserByToken(ctx)
		// if errToken != nil {
		// 	ctx.Error(errToken)
		// 	return
		// }

		err := s.DeleteQuestion(questionID, surveyID)
		if err != nil {
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}

		return
	})

	return r
}

// TitleRequest is the request structure for question POST api
type TitleRequest struct {
	Title string `json:"title" gorm:"column:title;type:varchar(255);not null;" example:"This is title"`
}

var (
	// ErrAccessDenied means user don't have access to this question
	ErrAccessDenied = errors.NewHTTPError(http.StatusForbidden, "Access to this question denied")
	// ErrBadRequest means params are not correct
	ErrBadRequest = errors.NewHTTPError(http.StatusBadRequest, "Bad request")
	// ErrInsertFailed means record is not persusted into table
	ErrInsertFailed = errors.NewHTTPError(http.StatusInternalServerError, "Insert record failed")
	// ErrNotFound means question was not found in the db
	ErrNotFound = errors.NewHTTPError(http.StatusNotFound, "Question not found")
	// ErrUpdateFailed means question was not updated
	ErrUpdateFailed = errors.NewHTTPError(http.StatusInternalServerError, "Question update failed")
	// ErrDeleteFailed means question was not deleted from the db
	ErrDeleteFailed = errors.NewHTTPError(http.StatusInternalServerError, "Question deletion failed")
)
