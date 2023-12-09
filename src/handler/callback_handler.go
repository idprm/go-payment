package handler

import (
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/services"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type CallbackHandler struct {
	cfg                *config.Secret
	rds                *redis.Client
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	orderService       services.IOrderService
	transactionService services.ITransactionService
	callbackService    services.ICallbackService
}

func NewCallbackHandler(
	cfg *config.Secret,
	rds *redis.Client,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	orderService services.IOrderService,
	transactionService services.ITransactionService,
	callbackService services.ICallbackService,
) *CallbackHandler {
	return &CallbackHandler{
		cfg:                cfg,
		rds:                rds,
		logger:             logger,
		zap:                zap,
		orderService:       orderService,
		transactionService: transactionService,
		callbackService:    callbackService,
	}
}

func (h *CallbackHandler) Midtrans() {
}

func (h *CallbackHandler) Nicepay() {
}

func (h *CallbackHandler) DragonPay() {
}

func (h *CallbackHandler) JazzCash() {
}

func (h *CallbackHandler) Momo() {
}

func (h *CallbackHandler) Razer() {
}
