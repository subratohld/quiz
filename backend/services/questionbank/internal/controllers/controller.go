package controllers

import (
	"context"
	cmn "github.com/subratohld/quiz/questionbank/internal/common/consts"
	"github.com/subratohld/quiz/questionbank/internal/services"
	"net/http"
)

import qpb "github.com/subratohld/quiz/questionbank/internal/qbproto"

type QuestionBank interface {
	CreateQuiz(context.Context, *qpb.CreateQuizRequest) (*qpb.CreateQuizResponse, error)
	AddQuestionsToQuizID(context.Context, *qpb.AddQuestionsToQuizRequest) (*qpb.AddQuestionsToQuizResponse, error)
	AddAnswerToQuestionID(ctx context.Context, request *qpb.AddAnswerToQuestionRequest) (*qpb.BaseResponse, error)
}

type questionBankController struct {
	svcManager *services.ServiceManager
}

func newQuestionBank(svcManager *services.ServiceManager) QuestionBank {
	return &questionBankController{
		svcManager: svcManager,
	}
}

func (q questionBankController) CreateQuiz(ctx context.Context, request *qpb.CreateQuizRequest) (*qpb.CreateQuizResponse, error) {
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

func (q questionBankController) AddQuestionsToQuizID(ctx context.Context, req *qpb.AddQuestionsToQuizRequest) (*qpb.AddQuestionsToQuizResponse, error) {
	authData := authModelMapper(req.GetAuthData())
	if authData.Error != nil {
		return addQuestionsResponseMapper(nil, http.StatusUnauthorized, cmn.ErrInAuthorization), nil
	}

	if len(req.GetQuestions()) == 0 {
		return addQuestionsResponseMapper(nil, http.StatusBadRequest, cmn.ErrInvalidReqForAddingQns), nil
	}

	if req.GetQuizId() == cmn.EmptyString {
		return addQuestionsResponseMapper(nil, http.StatusBadRequest, cmn.ErrInvalidQuizID), nil
	}

	successfullyAddedQns, err := q.svcManager.QBService.AddQuestionsToQuiz(ctx, addQuestionsReqMapper(authData.UserID, req.Questions))
	if err != nil {
		return addQuestionsResponseMapper(nil, http.StatusInternalServerError, err.Error()), nil
	}

	return addQuestionsResponseMapper(successfullyAddedQns, http.StatusOK, cmn.SuccessfulQuizCreation), nil
}

func (q questionBankController) AddAnswerToQuestionID(ctx context.Context, req *qpb.AddAnswerToQuestionRequest) (*qpb.BaseResponse, error) {
	return nil, nil
}
