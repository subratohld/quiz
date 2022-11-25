package repositories

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	cmnerr "github.com/subratohld/quiz/cmnlib/errors"
	"github.com/subratohld/quiz/cmnlib/logger"
	"github.com/subratohld/quiz/usermanagement/internal/common/consts"
	"github.com/subratohld/quiz/usermanagement/internal/databases"
	"github.com/subratohld/quiz/usermanagement/internal/models"
	"go.uber.org/zap"
)

type (
	Usermanagement interface {
		AddUser(ctx context.Context, user *models.User) (*models.User, error)
		GetUserByID(ctx context.Context, userId string) (*models.User, error)
	}

	usermanagement struct {
		DBManager *databases.DBManager
	}
)

var _ Usermanagement = (*usermanagement)(nil)

func newUsermanagement(dbManager *databases.DBManager) *usermanagement {
	return &usermanagement{
		DBManager: dbManager,
	}
}

func (u *usermanagement) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := u.DBManager.UsermanagementDB.ModelContext(ctx, user).Insert()
	if err != nil {
		logger.Logger().Error(consts.ErrAddUser, zap.Error(err))
		return nil, cmnerr.NewDBError(err.Error())
	}

	return user, nil
}

func (u *usermanagement) GetUserByID(ctx context.Context, userId string) (*models.User, error) {
	var user models.User
	err := u.DBManager.UsermanagementDB.ModelContext(ctx, &user).
		Where("user_id=?", userId).Select()

	if err != nil {
		logger.Logger().Error(fmt.Sprintf("%s '%s'", consts.ErrFindUser, userId), zap.Error(err))

		switch err {
		case pg.ErrNoRows:
			return nil, cmnerr.NewNotFoundError(err.Error())
		}

		return nil, cmnerr.NewDBError(err.Error())
	}

	return &user, nil
}
