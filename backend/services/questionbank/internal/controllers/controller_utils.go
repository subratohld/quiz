package controllers

import (
	"github.com/subratohld/quiz/cmnlib/errors"
	"github.com/subratohld/quiz/questionbank/internal/models"
	qpb "github.com/subratohld/quiz/questionbank/internal/qbproto"
)

func createQuizResponseMapper(quiz *models.Quiz, statusCode int32, message string) *qpb.CreateQuizResponse {
	response := &qpb.CreateQuizResponse{}

	if quiz != nil {
		response.Quiz = &qpb.Quiz{
			Id:          quiz.ID,
			Description: quiz.Description,
			CreatedBy:   quiz.CreatedBy,
			CreatedOn:   quiz.CreatedOn,
		}
	}

	response.Message = message
	response.StatusCode = statusCode

	return response
}

func createQuizRequestMapper(user string, desc string) *models.Quiz {
	return &models.Quiz{
		CreatedBy:   user,
		Description: desc,
	}
}

func authModelMapper(auth *qpb.AuthData) *models.Auth {
	authModel := &models.Auth{}
	if auth == nil {
		authModel.Error = errors.NewUnauthorizedError("No auth data provided")
	} else if len(auth.GetRoles()) > 0 || auth.GetUserId() == "" || auth.GetUserId() == "" {
		authModel.Error = errors.NewUnauthorizedError("No credentials data provided")
	} else {
		authModel.UserID = auth.GetUserId()
		authModel.OrgID = auth.GetOrgId()
		authModel.IsSystem = auth.GetIsSystem()
	}

	return authModel
}

func addQuestionsReqMapper(user string, reqQns []*qpb.Question) []*models.Question {
	var questions []*models.Question
	for _, qn := range reqQns {
		var options []*models.AnswerOptions
		for _, op := range qn.Options {
			options = append(options, &models.AnswerOptions{
				OptionId:          op.Id,
				OptionDescription: op.Description,
			})
		}
		questions = append(questions, &models.Question{
			Description: qn.Description,
			CreatedBy:   user,
			Options:     options,
		})
	}

	return questions
}

func addQuestionsResponseMapper(questions []*models.Question, statusCode int32, message string) *qpb.AddQuestionsToQuizResponse {
	response := &qpb.AddQuestionsToQuizResponse{}

	if len(questions) > 0 {
		for _, qn := range questions {

			var qpbOptions []*qpb.AnswerOptions
			if len(qn.Options) > 0 {
				for _, opt := range qn.Options {
					qpbOptions = append(qpbOptions, &qpb.AnswerOptions{
						Id:          opt.OptionId,
						Description: opt.OptionDescription,
					})
				}
			}

			response.Questions = append(response.Questions, &qpb.Question{
				Id:          qn.ID,
				Description: qn.Description,
				CreatedBy:   qn.CreatedBy,
				CreatedOn:   qn.CreatedOn,
				Options:     qpbOptions,
			})
		}
	}

	response.Message = message
	response.StatusCode = statusCode

	return response
}

func addCorrectAnsReqMapper(user string, answerOptions []string, linkedQnID string) *models.Answer {
	var ansOptionsModel []*models.AnswerOptions
	for _, answerOption := range answerOptions {
		ansOptionsModel = append(ansOptionsModel, &models.AnswerOptions{
			OptionId: answerOption,
		})
	}

	return &models.Answer{
		AnswerDetails: ansOptionsModel,
		CreatedBy:     user,
		LinkedQnId:    linkedQnID,
	}
}

func addCorrectAnsResMapper(answer *models.Answer, statusCode int32, message string) *qpb.AddCorrectAnswerResponse {
	response := &qpb.AddCorrectAnswerResponse{}

	if answer != nil {
		var listOfAnswerDetailsResp []*qpb.AnswerOptions
		for _, ansDetails := range answer.AnswerDetails {
			listOfAnswerDetailsResp = append(listOfAnswerDetailsResp, &qpb.AnswerOptions{
				Id:          ansDetails.OptionId,
				Description: ansDetails.OptionDescription,
			})
		}
		response.Answer = &qpb.Answer{
			Id:            answer.ID,
			AnswerDetails: listOfAnswerDetailsResp,
			CreatedBy:     answer.CreatedBy,
			CreatedOn:     answer.CreatedOn,
			UpdatedBy:     answer.UpdatedBy,
			UpdatedOn:     answer.UpdatedOn,
			LinkedQnId:    answer.LinkedQnId,
		}
	}

	response.Message = message
	response.StatusCode = statusCode

	return response
}
