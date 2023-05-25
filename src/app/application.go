package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/logger"
	"gorm.io/gorm"
)

type Application struct {
	cfg    *config.Secret
	db     *gorm.DB
	logger *logger.Logger
}

func NewApplication(
	cfg *config.Secret,
	db *gorm.DB,
	logger *logger.Logger,
) *Application {
	return &Application{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

type UrlMappings struct {
	cfg    *config.Secret
	db     *gorm.DB
	logger *logger.Logger
}

func NewUrlMappings(
	cfg *config.Secret,
	db *gorm.DB,
	logger *logger.Logger,
) *UrlMappings {
	return &UrlMappings{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

func (a *Application) Start() *fiber.App {
	urls := NewUrlMappings(a.cfg, a.db, a.logger)
	return urls.mapUrls()
}
