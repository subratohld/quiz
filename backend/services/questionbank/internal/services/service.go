package services

import (
	"context"

	"github.com/subratohld/quiz/questionbank/internal/models"
	"github.com/subratohld/quiz/questionbank/internal/repositories"
)

type (
	QuestionBank interface {
		CreateQuiz(ctx context.Context, quizModel *models.Quiz) (*models.Quiz, error)
		AddQuestionsToQuiz(ctx context.Context, questions []*models.Question) ([]*models.Question, error)
	}

	questionBankSvc struct {
		repoManager *repositories.RepoManager
	}
)

func newQuestionBankSvc(repoManager *repositories.RepoManager) QuestionBank {
	return &questionBankSvc{
		repoManager: repoManager,
	}
}

func (q *questionBankSvc) CreateQuiz(ctx context.Context, quizModel *models.Quiz) (*models.Quiz, error) {
	quizModel, err := q.repoManager.QBRepo.CreateQuiz(ctx, quizModel)
	if err != nil {
		return nil, err
	}

	return quizModel, nil
}

func (q *questionBankSvc) AddQuestionsToQuiz(ctx context.Context, questions []*models.Question) ([]*models.Question, error) {
	successfullyAddedQns, err := q.repoManager.QBRepo.AddQuestionsToQuiz(ctx, questions)
	if err != nil {
		return nil, err
	}

	return successfullyAddedQns, nil
}
