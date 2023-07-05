package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-payment/src/config"
	"github.com/idprm/go-payment/src/logger"
	"github.com/idprm/go-payment/src/services"
	"go.uber.org/zap"
)

type CallbackHandler struct {
	cfg                *config.Secret
	logger             *logger.Logger
	zap                *zap.SugaredLogger
	orderService       services.IOrderService
	transactionService services.ITransactionService
	callbackService    services.ICallbackService
}

func NewCallbackHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	zap *zap.SugaredLogger,
	orderService services.IOrderService,
	transactionService services.ITransactionService,
	callbackService services.ICallbackService,
) *CallbackHandler {
	return &CallbackHandler{
		cfg:                cfg,
		logger:             logger,
		zap:                zap,
		orderService:       orderService,
		transactionService: transactionService,
		callbackService:    callbackService,
	}
}

func (h *CallbackHandler) Razer(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
