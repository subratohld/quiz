package controllers

import (
	"context"
	"fmt"
	cmn "github.com/subratohld/quiz/questionbank/internal/common/consts"
	qpb "github.com/subratohld/quiz/questionbank/internal/qbproto"
	"github.com/subratohld/quiz/questionbank/internal/services"
	"net/http"
)

type QuestionBank interface {
	AddQuestions(context.Context, *qpb.AddQuestionsRequest) (*qpb.AddQuestionsResponse, error)
}

type questionBankController struct {
	svcManager *services.ServiceManager
}

func newQuestionBank(svcManager *services.ServiceManager) QuestionBank {
	return &questionBankController{
		svcManager: svcManager,
	}
}

// TODO: Add question type. MultipleChoiceQuestion(MCQ) / SingleChoiceQuestion(SCQ)
func (q questionBankController) AddQuestions(ctx context.Context, req *qpb.AddQuestionsRequest) (*qpb.AddQuestionsResponse, error) {
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

	return addQuestionsResponseMapper(successfullyAddedQns, http.StatusOK, fmt.Sprintf("%s %s", cmn.SuccessfulQuestionCreation, req.QuizId)), nil
}
