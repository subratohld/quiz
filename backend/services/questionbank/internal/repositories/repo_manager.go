package repositories

import (
	"github.com/subratohld/quiz/questionbank/internal/databases"
)

type RepoManager struct {
	QBRepo QuestionBank
}

func NewRepoManager(dbManager *databases.DBManager) *RepoManager {
	return &RepoManager{
		QBRepo: newQuestionBankRepo(dbManager),
	}
}
