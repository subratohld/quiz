package controllers

import (
	"github.com/subratohld/quiz/questionbank/internal/services"
)

type ControllerManager struct {
	Quiz
	QuestionBank
	CorrectAnswer
}

func NewControllerManager(svcManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		Quiz:          newQuiz(svcManager),
		QuestionBank:  newQuestionBank(svcManager),
		CorrectAnswer: newCorrectAnswer(svcManager),
	}
}
