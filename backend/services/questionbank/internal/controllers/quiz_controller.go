package controllers

import (
	"context"
	cmn "github.com/subratohld/quiz/questionbank/internal/common/consts"
	qpb "github.com/subratohld/quiz/questionbank/internal/qbproto"
	"github.com/subratohld/quiz/questionbank/internal/services"
	"net/http"
)

type Quiz interface {
	CreateQuiz(context.Context, *qpb.CreateQuizRequest) (*qpb.CreateQuizResponse, error)
}

type quizController struct {
	svcManager *services.ServiceManager
}

func newQuiz(svcManager *services.ServiceManager) Quiz {
	return &quizController{
		svcManager: svcManager,
	}
}

func (q quizController) CreateQuiz(ctx context.Context, request *qpb.CreateQuizRequest) (*qpb.CreateQuizResponse, error) {
	authData := authModelMapper(request.GetAuthData())
	if authData.Error != nil {
		return createQuizResponseMapper(nil, http.StatusUnauthorized, cmn.ErrInAuthorization), nil
	}

	quiz, err := q.svcManager.QBService.CreateQuiz(ctx, createQuizRequestMapper(authData.UserID, request.GetDescription()))
	if err != nil {
		return createQuizResponseMapper(nil, http.StatusInternalServerError, err.Error()), nil
	}

	return createQuizResponseMapper(quiz, http.StatusOK, cmn.SuccessfulQuizCreation), nil
}
