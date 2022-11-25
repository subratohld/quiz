package controllers

import (
	"github.com/subratohld/quiz/usermanagement/internal/services"
)

type ControllerManager struct {
	Usermanagement
}

func NewControllerManager(svcManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		Usermanagement: newUsermanagement(svcManager),
	}
}
