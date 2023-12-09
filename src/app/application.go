package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	cfg    *config.Secret
	db     *gorm.DB
	rds    *redis.Client
	logger *logger.Logger
	zap    *zap.SugaredLogger
	ctx    context.Context
}

func NewApplication(
	cfg *config.Secret,
	db *gorm.DB,
	rds *redis.Client,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	ctx context.Context,
) *Application {
	return &Application{
		cfg:    cfg,
		db:     db,
		rds:    rds,
		logger: logger,
		zap:    zap,
		ctx:    ctx,
	}
}

type UrlMappings struct {
	cfg    *config.Secret
	db     *gorm.DB
	rds    *redis.Client
	logger *logger.Logger
	zap    *zap.SugaredLogger
	ctx    context.Context
}

func NewUrlMappings(
	cfg *config.Secret,
	db *gorm.DB,
	rds *redis.Client,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	ctx context.Context,
) *UrlMappings {
	return &UrlMappings{
		cfg:    cfg,
		db:     db,
		rds:    rds,
		logger: logger,
		zap:    zap,
		ctx:    ctx,
	}
}

func (a *Application) Start() *fiber.App {
	urls := NewUrlMappings(a.cfg, a.db, a.rds, a.logger, a.zap, a.ctx)
	return urls.mapUrls()
}
