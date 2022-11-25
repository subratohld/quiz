package services

import (
	"context"

	"github.com/subratohld/quiz/usermanagement/internal/models"
	"github.com/subratohld/quiz/usermanagement/internal/repositories"
)

type (
	Usermanagement interface {
		AddUser(ctx context.Context, user *models.User) (*models.User, error)
		GetUserByID(ctx context.Context, userId string) (*models.User, error)
	}

	usermanagement struct {
		repoManager *repositories.RepoManager
	}
)

var _ Usermanagement = (*usermanagement)(nil)

func newUsermanagement(repoManager *repositories.RepoManager) *usermanagement {
	return &usermanagement{
		repoManager: repoManager,
	}
}

func (u *usermanagement) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	return u.repoManager.UMRepo.AddUser(ctx, user)
}

func (u *usermanagement) GetUserByID(ctx context.Context, userId string) (*models.User, error) {
	return u.repoManager.UMRepo.GetUserByID(ctx, userId)
}
