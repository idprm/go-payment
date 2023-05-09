package db

import (
	"github.com/idprm/go-payment/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(conf *config.Secret) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(conf.Db.Source), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// DEBUG ON CONSOLE
	// db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate()
	if err != nil {
		return nil, err
	}

	return db, nil
}
