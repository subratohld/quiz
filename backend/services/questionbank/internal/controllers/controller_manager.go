package controllers

import (
	"github.com/subratohld/quiz/questionbank/internal/services"
)

type ControllerManager struct {
	QuestionBank
}

func NewControllerManager(svcManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		QuestionBank: newQuestionBank(svcManager),
	}
}
