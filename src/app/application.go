package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"gorm.io/gorm"
)

type Application struct {
	cfg *config.Secret
	db  *gorm.DB
}

func NewApplication(
	cfg *config.Secret,
	db *gorm.DB,
) *Application {
	return &Application{
		cfg: cfg,
		db:  db,
	}
}

type UrlMappings struct {
	cfg *config.Secret
	db  *gorm.DB
}

func NewUrlMappings(
	cfg *config.Secret,
	db *gorm.DB,
) *UrlMappings {
	return &UrlMappings{
		cfg: cfg,
		db:  db,
	}
}

func (a *Application) Start() *fiber.App {
	urls := NewUrlMappings(a.cfg, a.db)
	return urls.mapUrls()
}
