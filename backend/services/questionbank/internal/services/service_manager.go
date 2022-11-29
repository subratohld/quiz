package services

import (
	"github.com/subratohld/quiz/questionbank/internal/repositories"
)

type ServiceManager struct {
	QBService QuestionBank
}

func NewServiceManager(repoManager *repositories.RepoManager) *ServiceManager {
	return &ServiceManager{
		QBService: newQuestionBankSvc(repoManager),
	}
}
