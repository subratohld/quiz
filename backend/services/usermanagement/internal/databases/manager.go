package databases

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/subratohld/quiz/cmnlib/logger"
	"github.com/subratohld/quiz/usermanagement/internal/common/config"
	"go.uber.org/zap"
)

type DBManager struct {
	UsermanagementDB orm.DB
}

func NewDBManager(conf *config.Config) *DBManager {
	dbManager := new(DBManager)

	if err := dbManager.init(conf); err != nil {
		return nil
	}

	return dbManager
}

func (dbManager *DBManager) init(conf *config.Config) error {
	db := pg.Connect(&pg.Options{
		Addr:     conf.Databases.Postgres.Host,
		Database: conf.Databases.Postgres.DBName,
		User:     conf.Databases.Postgres.Secret.Username,
		Password: conf.Databases.Postgres.Secret.Password,
	})

	if err := db.Ping(context.TODO()); err != nil {
		logger.Logger().Error("could not connect to postgres db", zap.Error(err))
		return err
	}

	dbManager.UsermanagementDB = db

	return nil
}
