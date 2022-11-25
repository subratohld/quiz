package repositories

import (
	"github.com/subratohld/quiz/usermanagement/internal/databases"
)

type RepoManager struct {
	UMRepo Usermanagement
}

func NewRepoManager(dbManager *databases.DBManager) *RepoManager {
	return &RepoManager{
		UMRepo: newUsermanagement(dbManager),
	}
}
