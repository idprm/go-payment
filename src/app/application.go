package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	APP_TZ   string = utils.GetEnv("APP_TZ")
	LOG_PATH string = utils.GetEnv("LOG_PATH")
)

type Application struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *logger.Logger
	zap    *zap.SugaredLogger
	ctx    context.Context
}

func NewApplication(
	db *gorm.DB,
	rds *redis.Client,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	ctx context.Context,
) *Application {
	return &Application{
		db:     db,
		rds:    rds,
		logger: logger,
		zap:    zap,
		ctx:    ctx,
	}
}

type UrlMappings struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *logger.Logger
	zap    *zap.SugaredLogger
	ctx    context.Context
}

func NewUrlMappings(
	db *gorm.DB,
	rds *redis.Client,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	ctx context.Context,
) *UrlMappings {
	return &UrlMappings{
		db:     db,
		rds:    rds,
		logger: logger,
		zap:    zap,
		ctx:    ctx,
	}
}

func (a *Application) Start() *fiber.App {
	urls := NewUrlMappings(a.db, a.rds, a.logger, a.zap, a.ctx)
	return urls.mapUrls()
}
