package controllers

import (
	"context"
	"fmt"
	cmn "github.com/subratohld/quiz/questionbank/internal/common/consts"
	qpb "github.com/subratohld/quiz/questionbank/internal/qbproto"
	"github.com/subratohld/quiz/questionbank/internal/services"
	"net/http"
)

type CorrectAnswer interface {
	AddCorrectAnswer(ctx context.Context, request *qpb.AddCorrectAnswerRequest) (*qpb.AddCorrectAnswerResponse, error)
}

type correctAnswerController struct {
	svcManager *services.ServiceManager
}

func newCorrectAnswer(svcManager *services.ServiceManager) CorrectAnswer {
	return &correctAnswerController{
		svcManager: svcManager,
	}
}

// TODO : While processing the request to add correct answer, check if the linked question type is MCQ / SCQ . User has to provide the answer details accordingly.
func (c correctAnswerController) AddCorrectAnswer(ctx context.Context, req *qpb.AddCorrectAnswerRequest) (*qpb.AddCorrectAnswerResponse, error) {
	authData := authModelMapper(req.GetAuthData())
	if authData.Error != nil {
		return addCorrectAnsResMapper(nil, http.StatusUnauthorized, cmn.ErrInAuthorization), nil
	}

	if len(req.GetAnswerDetails()) == 0 {
		return addCorrectAnsResMapper(nil, http.StatusBadRequest, cmn.ErrAddingAnswerDetails), nil
	}

	if req.GetLinkedQnId() == cmn.EmptyString {
		return addCorrectAnsResMapper(nil, http.StatusBadRequest, cmn.ErrLinkedQuestionIDMissing), nil
	}

	answer, err := c.svcManager.QBService.AddCorrectAnswerForQuestionID(ctx, addCorrectAnsReqMapper(authData.UserID, req.GetAnswerDetails(), req.LinkedQnId))
	if err != nil {
		return addCorrectAnsResMapper(nil, http.StatusInternalServerError, err.Error()), nil
	}

	return addCorrectAnsResMapper(answer, http.StatusOK, fmt.Sprintf("%s %s", cmn.SuccessfulCorrectAnsCreation, req.LinkedQnId)), nil
}
