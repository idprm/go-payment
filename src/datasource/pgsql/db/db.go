package db

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPgSQL(conf *config.Secret) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(conf.Db.SourcePgSql), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(
		&entity.Application{},
		&entity.Gateway{},
		&entity.Channel{},
		&entity.Order{},
		&entity.Transaction{},
		&entity.Payment{},
		&entity.Callback{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}