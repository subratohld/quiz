package services

import (
	"github.com/subratohld/quiz/usermanagement/internal/repositories"
)

type ServiceManager struct {
	UMService Usermanagement
}

func NewServiceManager(repoManager *repositories.RepoManager) *ServiceManager {
	return &ServiceManager{
		UMService: newUsermanagement(repoManager),
	}
}
