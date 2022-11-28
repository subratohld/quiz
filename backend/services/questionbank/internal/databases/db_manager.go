package databases

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/subratohld/quiz/cmnlib/logger"
	"github.com/subratohld/quiz/questionbank/internal/common/config"
)

type DBManager struct {
	QuestionBankDB orm.DB
}

func NewDBManager(conf *config.Config) *DBManager {
	dbManager := new(DBManager)

	if err := dbManager.init(conf); err != nil {
		return nil
	}

	return dbManager
}

func (dbManager *DBManager) init(conf *config.Config) error {
	pgDb := pg.Connect(&pg.Options{
		Addr:     conf.Databases.Postgres.Host,
		Database: conf.Databases.Postgres.DBName,
		User:     conf.Databases.Postgres.Secret.Username,
		Password: conf.Databases.Postgres.Secret.Password,
	})

	if err := pgDb.Ping(context.Background()); err != nil {
		logger.Logger().Sugar().Errorf("could not connect to postgres db. Err: %v", err)
		return err
	}

	dbManager.QuestionBankDB = pgDb

	return nil
}
