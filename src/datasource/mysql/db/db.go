package db

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySQL(conf *config.Secret) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(conf.Db.SourceMySql), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(
		&entity.Application{},
		&entity.Country{},
		&entity.Gateway{},
		&entity.Channel{},
		&entity.Order{},
		&entity.Payment{},
		&entity.Callback{},
		&entity.Refund{},
		&entity.Transaction{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
