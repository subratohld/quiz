package repositories

import (
	"context"
	cmnErr "github.com/subratohld/quiz/cmnlib/errors"
	"github.com/subratohld/quiz/cmnlib/logger"
	cmnutil "github.com/subratohld/quiz/cmnlib/util"
	"github.com/subratohld/quiz/questionbank/internal/common/consts"
	"github.com/subratohld/quiz/questionbank/internal/databases"
	"github.com/subratohld/quiz/questionbank/internal/models"
	"go.uber.org/zap"
	"time"
)

type (
	QuestionBank interface {
		CreateQuiz(ctx context.Context, quizModel *models.Quiz) (*models.Quiz, error)
		AddQuestionsToQuiz(ctx context.Context, questions []*models.Question) ([]*models.Question, error)
	}

	questionBankRepo struct {
		DBManager *databases.DBManager
	}
)

func newQuestionBankRepo(dbManager *databases.DBManager) QuestionBank {
	return &questionBankRepo{
		DBManager: dbManager,
	}
}

func (q *questionBankRepo) CreateQuiz(ctx context.Context, quizModel *models.Quiz) (*models.Quiz, error) {
	quizModel.ID = cmnutil.GenerateUUID()
	quizModel.CreatedOn = time.Now().Unix()
	_, err := q.DBManager.QuestionBankDB.ModelContext(ctx, quizModel).Insert()
	if err != nil {
		logger.Logger().Error(consts.ErrAddingQuizInDB, zap.Error(err))
		return nil, cmnErr.NewDBError(err.Error())
	}

	return quizModel, nil
}

func (q *questionBankRepo) AddQuestionsToQuiz(ctx context.Context, questions []*models.Question) ([]*models.Question, error) {
	var successfulInsertionQnModels []*models.Question
	for _, question := range questions {
		question.ID = cmnutil.GenerateUUID()
		question.CreatedOn = time.Now().Unix()
		_, err := q.DBManager.QuestionBankDB.ModelContext(ctx, question).Insert()

		if err != nil {
			logger.Logger().Error(consts.ErrAddingQuestionInDB, zap.Error(err))
			return nil, cmnErr.NewDBError(err.Error())
		}

		successfulInsertionQnModels = append(successfulInsertionQnModels, question)
	}

	return successfulInsertionQnModels, nil
}
